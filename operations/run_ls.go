package operations

import (
	"github.com/matsu-chara/gol/kvs"
)

// RunLs run ls
func RunLs(filepath string) ([]kvs.Entry, error) {
	db, err := kvs.Open(filepath)
	defer func() {
		db.Close()
	}()
	if err != nil {
		return nil, err
	}

	return db.List(), nil
}
