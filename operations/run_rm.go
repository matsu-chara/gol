package operations

import "github.com/matsu-chara/gol/kvs"

// RunRm run rm
func RunRm(filepath string, key string, registeredBy string) error {
	db, err := kvs.Open(filepath, kvs.ReadAndWrite)
	defer db.Close()
	if err != nil {
		return err
	}

	err = db.Remove(key, registeredBy)
	if err != nil {
		return err
	}
	return db.Save()
}
