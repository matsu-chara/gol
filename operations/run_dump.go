package operations

import "github.com/matsu-chara/gol/kvs"

// RunDump run dump
func RunDump(filepath string) (map[string]kvs.Value, error) {
	db, err := kvs.Open(filepath, kvs.ReadOnly)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	return db.Dump(), nil
}
