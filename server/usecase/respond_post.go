package usecase

import (
	"net/http"

	"github.com/matsu-chara/gol/operations"
)

// Post adds new key value
func Post(filepath string, key string, value string, w http.ResponseWriter, r *http.Request) {
	entry, err := operations.RunGet(filepath, key)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	if entry != nil {
		status := http.StatusConflict
		http.Error(w, http.StatusText(status), status)
		return
	}

	err = operations.RunAdd(filepath, key, value)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	return
}
