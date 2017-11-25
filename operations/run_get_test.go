package operations

import (
	"os"
	"testing"
)

func TestRunGet(t *testing.T) {
	testFile := tempTest("get")
	defer os.Remove(testFile)

	err := initDb(testFile)
	if err != nil {
		t.Errorf("db init error %s", err)
	}

	result, err := RunGet(testFile, "k1")
	if err != nil {
		t.Errorf("get error %s", err)
	}
	if result.Value.Link != "v1" {
		t.Errorf("unexpected value %s", result)
	}
}
