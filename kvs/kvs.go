package kvs

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

// KVS which have data and metadata
type KVS struct {
	filename string
	Data
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
	newJSON, err := json.Marshal(kvs.data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(kvs.filename, newJSON, os.FileMode(filePermission)); err != nil {
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
		if err := ioutil.WriteFile(filename, []byte("{}"), os.FileMode(filePermission)); err != nil {
			return err
		}
	}
	return nil
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
