package operations

import (
	"github.com/matsu-chara/gol/kvs"
)

// RunGet run get
func RunGet(filepath string, key string) (*kvs.Entry, error) {
	db, err := kvs.Open(filepath)
	defer func() {
		db.Close()
	}()
	if err != nil {
		return nil, err
	}

	entry, found := db.Get(key)
	if !found {
		return nil, nil
	}

	return entry, nil
}
