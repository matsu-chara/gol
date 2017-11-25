package operations

import (
	"os"
	"path/filepath"

	"github.com/matsu-chara/gol/kvs"
)

func tempTest(name string) string {
	return filepath.Join(os.TempDir(), "/gol_test_"+name)
}

func initDb(testFile string) error {
	db, err := kvs.Open(testFile)
	if err != nil {
		return err
	}
	entry, err := kvs.NewEntry("k1", kvs.Value{Link: "v1"})
	if err != nil {
		return err
	}
	err = db.Put(entry)
	if err != nil {
		return err
	}
	entry2, err := kvs.NewEntry("k2", kvs.Value{Link: "v2"})
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
