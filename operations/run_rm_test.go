package operations

import (
	"github.com/matsu-chara/gol/kvs"
	"os"
	"testing"
)

func TestRunRm(t *testing.T) {
	testFile := tempTest("add")
	defer os.Remove(testFile)
	initDb(testFile)

	RunRm(testFile, "k1")

	db, err := kvs.Open(testFile)
	if err != nil {
		t.Errorf("kvs open failed %s", err)
	}
	result, isExists := db.Get("k1")
	if isExists || result != nil {
		t.Errorf("key_add was found. result = %s", result)
	}
}
