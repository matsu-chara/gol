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

	kvs, err := Open(testFile, ReadOnly)
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

	kvs, err := Open(testFile, ReadAndWrite)
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
	kvs.Close()

	kvs2, err4 := Open(testFile, ReadAndWrite)
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
	kvs2.Close()
}

func TestClose(t *testing.T) {
	dir := os.TempDir()
	testFile := filepath.Join(dir, "gol_test_close")
	defer os.Remove(testFile)

	kvs, _ := Open(testFile, ReadOnly)
	kvs.Close() // no panic
}

func TestCache(t *testing.T) {
	dir := os.TempDir()
	testFile := filepath.Join(dir, "gol_test_cache")
	defer os.Remove(testFile)

	kvs, err := Open(testFile, ReadAndWrite)
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
	kvs.Close()

	kvs2, err4 := Open(testFile, ReadOnly)
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
	kvs2.Close()

	os.Remove(testFile) // file will be removed. but in-memory cache will remain.
	kvs3, err5 := Open(testFile, ReadOnly)
	if err5 != nil {
		t.Errorf("open failed %s", err5)
	}

	entry2, _ := kvs3.Get("k_test") // returning from cache
	if (*entry2 != Entry{Key: "k_test", Value: Value{Link: "v1"}}) {
		t.Errorf("unexpected entry %s", entry2)
	}
	kvs3.Close()
}
