package operations

import (
	"os"
	"testing"

	"github.com/matsu-chara/gol/kvs"
)

func TestRunAdd(t *testing.T) {
	testFile := tempTest("add")
	registeredBy := "test"
	isForce := false
	defer os.Remove(testFile)
	initDb(testFile)

	RunAdd(testFile, "key_add", "value", registeredBy, isForce)

	db, err := kvs.Open(testFile, kvs.ReadAndWrite)
	defer db.Close()
	if err != nil {
		t.Errorf("kvs open failed %s", err)
	}
	result, isExists := db.Get("key_add")
	if !isExists {
		t.Error("key_add was not found.")
	}

	if result.Value.Link != "value" {
		t.Errorf("%s is not expected key key_add", result)
	}
	if result.Value.RegisteredBy != registeredBy {
		t.Errorf("%s is not expected key key_add", result)
	}
}

func TestRunAddConflict(t *testing.T) {
	testFile := tempTest("add")
	isForce := false
	defer os.Remove(testFile)
	initDb(testFile)

	err1 := RunAdd(testFile, "key_add", "value", "", isForce)
	if err1 != nil {
		t.Errorf("failed to first time add. %s", err1)
	}

	err2 := RunAdd(testFile, "key_add", "value", "", isForce)
	if err2 == nil {
		t.Errorf("add twice should be err. but err == nil")
	}
}

func TestRunAddForce(t *testing.T) {
	testFile := tempTest("add")
	isForce := true
	defer os.Remove(testFile)
	initDb(testFile)

	err1 := RunAdd(testFile, "key_add", "value", "", isForce)
	if err1 != nil {
		t.Errorf("failed to first time add. %s", err1)
	}

	err2 := RunAdd(testFile, "key_add", "value2", "", isForce)
	if err2 != nil {
		t.Errorf("add twice should replace value. %s", err2)
	}

	v, err := RunGet(testFile, "key_add")
	if v.Value.Link != "value2" || err != nil {
		t.Errorf("value wasn't replaced or err was occurred. %s %s", v.Value.Link, err)
	}

	err3 := RunAdd(testFile, "key_add", "value3", "foo", isForce)
	if err3 == nil {
		t.Errorf("add no same registeredBy should not replace value. %s", err3)
	}

	v, err = RunGet(testFile, "key_add")
	if v.Value.Link != "value2" || err != nil {
		t.Errorf("value was replaced or err was occurred. %s %s", v.Value.Link, err)
	}
}
