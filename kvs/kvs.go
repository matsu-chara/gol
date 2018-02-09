package kvs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

// KVS which have data and metadata
type KVS struct {
	filename string
	*Data
}

const (
	dirPermission  = 0755
	filePermission = 0644
)

// for inmemory cache.
// this will shared by only same process. so, CLI will be unaffected by this cache & lock.
var inMemoryCacheLock sync.RWMutex
var inMemoryCache *Data

// Open returns KVS data and metadata.
func Open(filename string, permission Permission) (*KVS, error) {
	switch permission {
	case ReadOnly:
		inMemoryCacheLock.RLock() // prevent read file, during writing file. will release in Close method
		if inMemoryCache != nil {
			kvs := KVS{
				filename,
				inMemoryCache,
			}
			return &kvs, nil // return cache
		}
	case ReadAndWrite:
		inMemoryCacheLock.Lock() // will release in Close method
		inMemoryCache = nil
	}
	if err := prepareFile(filename); err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var values map[string]interface{}
	err = json.Unmarshal(raw, &values)
	if err != nil {
		return nil, err
	}

	data, err := toData(values, permission)
	if err != nil {
		return nil, err
	}

	kvs := KVS{
		filename: filename,
		Data:     &data,
	}
	return &kvs, nil
}

// Save all data.
func (kvs *KVS) Save() error {
	if kvs.Data.permission != ReadAndWrite {
		return errors.New("save error. kvs was opend with ReadOnly Permission")
	}

	newJSON, err := json.Marshal(kvs.data)
	if err != nil {
		return err
	}

	err = fileWrite(kvs.filename, newJSON)
	if err != nil {
		return err
	}

	// update inmemory cache
	inMemoryCache = &Data{
		data:       kvs.Data.data,
		permission: ReadOnly, // as read only kvs
	}
	return nil
}

// Close KVS (currently do nothing)
func (kvs *KVS) Close() {
	switch kvs.Data.permission {
	case ReadOnly:
		inMemoryCacheLock.RUnlock()
	case ReadAndWrite:
		inMemoryCacheLock.Unlock()
	}
	kvs.Data = nil
	return
}

func toData(compatibleData map[string]interface{}, permission Permission) (Data, error) {
	data := map[string]Value{}

	for key, compatibleValue := range compatibleData {
		var value *Value

		switch compatibleValue.(type) {
		case string: // for ~v0.3.0
			value = &Value{
				Link: compatibleValue.(string),
			}
		case map[string]interface{}: // for v0.4.0~
			var createdAt time.Time
			err := createdAt.UnmarshalText([]byte(compatibleValue.(map[string]interface{})["CreatedAt"].(string)))
			if err != nil {
				return Data{}, err
			}
			value = &Value{
				Link:         compatibleValue.(map[string]interface{})["Link"].(string),
				RegisteredBy: compatibleValue.(map[string]interface{})["RegisteredBy"].(string),
				CreatedAt:    createdAt,
			}
		default:
			return Data{}, errors.New("can't decode data at key: " + key)
		}
		data[key] = *value
	}

	return Data{
		data:       data,
		permission: permission,
	}, nil
}

func prepareFile(filename string) error {
	if err := os.MkdirAll(path.Dir(filename), dirPermission); err != nil {
		return err
	}
	if !isExist(filename) {
		err := fileWrite(filename, []byte("{}"))
		if err != nil {
			return err
		}
	}
	return nil
}

func fileWrite(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, filePermission)
	defer f.Close()

	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return nil
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
