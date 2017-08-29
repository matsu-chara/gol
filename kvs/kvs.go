package kvs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

// KVS which have data and metadata
type KVS struct {
	filename string
	Data
	mutex sync.Mutex
}

const (
	dirPermission  = 0755
	filePermission = 0644
)

// Open returns KVS data and metadata.
func Open(filename string) (*KVS, error) {
	if err := prepareFile(filename); err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var values map[string]string
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil, err
	}

	kvs := KVS{
		filename: filename,
		Data: Data{
			data: values,
		},
	}
	return &kvs, nil
}

// Save all data.
func (kvs *KVS) Save() error {
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()

	newJSON, err := json.Marshal(kvs.data)
	if err != nil {
		return err
	}

	err = fileWrite(kvs.filename, newJSON)
	if err != nil {
		return err
	}

	kvs.Close()
	return nil
}

// Close KVS (currently do nothing)
func (kvs *KVS) Close() {
	// do nothing
	return
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
	n, err := f.Write(data)
	if err != nil {
		return err
	}
	if err == nil && n < len(data) {
		return fmt.Errorf("Short Write. expected=%d actual=%d", len(data), n)
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
