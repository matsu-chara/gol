package operations

import (
	"github.com/matsu-chara/gol/kvs"
	"time"
)

// RunAdd run add
func RunAdd(filepath string, key string, link string, registeredBy string, isForce bool) error {
	value := kvs.Value{
		Link:         link,
		RegisteredBy: registeredBy,
		CreatedAt:    time.Now(),
	}
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
		err = db.Remove(key, registeredBy)
		if err != nil {
			return err
		}
	}

	if err := db.Put(entry); err != nil {
		return err
	}

	return db.Save()
}
