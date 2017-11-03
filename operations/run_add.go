package operations

import (
	"github.com/matsu-chara/gol/kvs"
)

// RunAdd run add
func RunAdd(filepath string, key string, value string, isForce bool) error {
	entry, err := kvs.NewEntry(key, value)
	if err != nil {
		return err
	}

	db, err := kvs.Open(filepath)
	defer func() {
		db.Close()
	}()
	if err != nil {
		return err
	}

	if isForce {
		db.Remove(key)
	}

	if err := db.Put(entry); err != nil {
		return err
	}

	if err := db.Save(); err != nil {
		return err
	}

	return nil
}
