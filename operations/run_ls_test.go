package operations

import (
	"os"
	"testing"
)

func TestRunLs(t *testing.T) {
	testFile := tempTest("dump")
	defer os.Remove(testFile)

	err := initDb(testFile)
	if err != nil {
		t.Errorf("db init error %s", err)
	}
	result, err := RunLs(testFile)
	if err != nil {
		t.Errorf("dump error %s", err)
	}

	if len(result) != 2 {
		t.Errorf("%s has unexpected length", result)
	}

	if result[0].Key != "k1" || result[1].Key != "k2" {
		t.Errorf("unexpected result %s", result)
	}
}
