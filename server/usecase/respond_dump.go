package usecase

import (
	"encoding/json"
	"net/http"

	"github.com/matsu-chara/gol/operations"
)

// Dump dumps all links in kvs
func Dump(filepath string, w http.ResponseWriter) {
	dumped, err := operations.RunDump(filepath)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	dumpedJSON, err := json.MarshalIndent(dumped, "", "\t")
	if err != nil {
		respondInternalServerError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dumpedJSON)
	return
}
