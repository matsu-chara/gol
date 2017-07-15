package use_case

import (
	"encoding/json"
	"fmt"
	"github.com/matsu-chara/gol/operations"
	"net/http"
)

func Dump(filepath string, w http.ResponseWriter) {
	dumped, err := operations.RunDump(filepath)
	if err != nil {
		fmt.Println(err)
		respondInternalServerError(w)
		return
	}
	dumpedJSON, err := json.MarshalIndent(dumped, "", "\t")
	if err != nil {
		fmt.Println(err)
		respondInternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(dumpedJSON)
	return
}
