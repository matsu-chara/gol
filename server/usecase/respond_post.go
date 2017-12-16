package usecase

import (
	"net/http"

	"github.com/matsu-chara/gol/operations"
)

// Post adds ew key value
func Post(filepath string, key string, link string, registeredBy string, isForce bool, w http.ResponseWriter, r *http.Request) {
	entry, err := operations.RunGet(filepath, key)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}

	if entry != nil {
		if isForce {
			if !entry.Value.IsRegisteredBy(registeredBy) {
				respondAsNotRegisteredBy(registeredBy, entry, w)
				return
			} // else will delete & add entry
		} else {
			status := http.StatusConflict
			http.Error(w, http.StatusText(status), status)
			return
		}
	}

	err = operations.RunAdd(filepath, key, link, registeredBy, isForce)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	return
}
