package operations

import (
	"os"
	"testing"

	"github.com/matsu-chara/gol/kvs"
)

func TestRunRm(t *testing.T) {
	testFile := tempTest("add")
	defer os.Remove(testFile)
	initDb(testFile)

	RunRm(testFile, "k1", "")

	db, err := kvs.Open(testFile)
	if err != nil {
		t.Errorf("kvs open failed %s", err)
	}
	result, isExists := db.Get("k1")
	if isExists || result != nil {
		t.Errorf("k1 was found. result = %s", result)
	}
}

func TestRunRmSameRegisteredBy(t *testing.T) {
	registeredBy := "some_body"
	isForce := false

	testFile := tempTest("add")
	defer os.Remove(testFile)

	err1 := RunAdd(testFile, "k1", "value", registeredBy, isForce)
	if err1 != nil {
		t.Errorf("failed to first time add. %s", err1)
	}

	RunRm(testFile, "k1", registeredBy)

	db, err := kvs.Open(testFile)
	if err != nil {
		t.Errorf("kvs open failed %s", err)
	}
	result, isExists := db.Get("k1")
	if isExists || result != nil {
		t.Errorf("k1 was found. result = %s", result)
	}
}

func TestRunRmDifferentRegisteredBy(t *testing.T) {
	registeredBy := "some_body"
	isForce := false

	testFile := tempTest("add")
	defer os.Remove(testFile)

	err1 := RunAdd(testFile, "k1", "value", registeredBy, isForce)
	if err1 != nil {
		t.Errorf("failed to first time add. %s", err1)
	}

	RunRm(testFile, "k1", "")

	db, err := kvs.Open(testFile)
	if err != nil {
		t.Errorf("kvs open failed %s", err)
	}
	result, isExists := db.Get("k1")
	if !isExists || result == nil {
		t.Errorf("k1 was not found. result = %s", result)
	}
}
