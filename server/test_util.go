package server

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/matsu-chara/gol/kvs"
)

func tempTest(name string) string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("/gol_test_%s_%d", name, time.Now().UnixNano()))
}

func initDb(testFile string) error {
	db, err := kvs.Open(testFile, kvs.ReadAndWrite)
	defer db.Close()
	if err != nil {
		return err
	}

	entry, err := kvs.NewEntry("k1", kvs.Value{Link: "http://test/v1"})
	if err != nil {
		return err
	}
	err = db.Put(entry)
	if err != nil {
		return err
	}
	entry2, err := kvs.NewEntry("k2", kvs.Value{Link: "http://v2"})
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
