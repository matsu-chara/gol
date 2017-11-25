package kvs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpen(t *testing.T) {
	dir := os.TempDir()
	testFile := filepath.Join(dir, "gol_test_open")
	defer os.Remove(testFile)

	kvs, err := Open(testFile)
	defer kvs.Close()
	if err != nil {
		t.Errorf("open kvs failed %s", err)
	}

	if kvs == nil {
		t.Error("kvs was nil")
	}
}

func TestSave(t *testing.T) {
	dir := os.TempDir()
	testFile := filepath.Join(dir, "gol_test_save")
	defer os.Remove(testFile)

	kvs, err := Open(testFile)
	defer kvs.Close()
	if err != nil {
		t.Errorf("open kvs failed %s", err)
	}

	if kvs == nil {
		t.Error("kvs was nil")
	}

	err2 := kvs.Put(&Entry{Key: "k_test", Value: Value{Link: "v1"}})
	if err2 != nil {
		t.Errorf("put failed %s", err2)
	}

	err3 := kvs.Save()
	if err2 != nil {
		t.Errorf("save failed %s", err3)
	}

	kvs2, err4 := Open(testFile)
	defer kvs2.Close()
	if err4 != nil {
		t.Errorf("save failed %s", err4)
	}

	entry, isExist := kvs2.Get("k_test")
	if !isExist {
		t.Error("save was ignored")
	}
	if (*entry != Entry{Key: "k_test", Value: Value{Link: "v1"}}) {
		t.Errorf("unexpected entry %s", entry)
	}
}

func TestClose(t *testing.T) {
	dir := os.TempDir()
	testFile := filepath.Join(dir, "gol_test_close")
	defer os.Remove(testFile)

	kvs, _ := Open(testFile)
	kvs.Close() // no panic
}
