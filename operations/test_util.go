package operations

import (
	"github.com/matsu-chara/gol/kvs"
	"os"
)

func tempTest(name string) string {
	return os.TempDir() + "gol_test_" + name
}

func initDb(testFile string) error {
	db, err := kvs.Open(testFile)
	if err != nil {
		return err
	}
	entry, err := kvs.NewEntry("k1", "v1")
	if err != nil {
		return err
	}
	err = db.Put(entry)
	if err != nil {
		return err
	}
	entry2, err := kvs.NewEntry("k2", "v2")
	if err != nil {
		return err
	}
	err = db.Put(entry2)
	if err != nil {
		return err
	}
	err = db.Save()
	if err != nil {
		return err
	}
	return nil
}
