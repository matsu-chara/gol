package operations

import (
	"os"
	"reflect"
	"testing"
)

func TestRunDump(t *testing.T) {
	testFile := tempTest("dump")
	defer os.Remove(testFile)

	err := initDb(testFile)
	if err != nil {
		t.Errorf("db init error %s", err)
	}
	result, err := RunDump(testFile)
	if err != nil {
		t.Errorf("dump error %s", err)
	}

	if !reflect.DeepEqual(result, map[string]string{"k1": "v1", "k2": "v2"}) {
		t.Errorf("unexpected result %s", result)
	}
}
