package operations

import "github.com/matsu-chara/gol/kvs"

// RunRm run rm
func RunRm(filepath string, key string) error {
	db, err := kvs.Open(filepath)
	defer func() {
		db.Close()
	}()
	if err != nil {
		return err
	}

	db.Remove(key)
	return db.Save()
}
