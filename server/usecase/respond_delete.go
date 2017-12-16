package usecase

import (
	"github.com/matsu-chara/gol/operations"
	"net/http"
)

// Delete remove key
func Delete(filepath string, key string, registeredBy string, w http.ResponseWriter, r *http.Request) {
	entry, err := operations.RunGet(filepath, key)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	if !entry.Value.IsRegisteredBy(registeredBy) {
		respondAsNotRegisteredBy(registeredBy, entry, w)
		return
	}

	err = operations.RunRm(filepath, key, registeredBy)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	return
}
