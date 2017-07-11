package operations

import (
	"github.com/matsu-chara/gol/kvs"
	"os"
	"testing"
)

func TestRunAdd(t *testing.T) {
	testFile := tempTest("add")
	defer os.Remove(testFile)
	initDb(testFile)

	RunAdd(testFile, "key_add", "value")

	db, err := kvs.Open(testFile)
	if err != nil {
		t.Errorf("kvs open failed %s", err)
	}
	result, isExists := db.Get("key_add")
	if !isExists {
		t.Error("key_add was not found.")
	}

	if result.Value != "value" {
		t.Errorf("%s is not expected key key_add", result)
	}
}
