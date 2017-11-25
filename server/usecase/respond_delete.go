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
	if entry.Value.RegisteredBy != registeredBy {
		contactInfoString := "please contact to " + entry.Value.RegisteredBy
		if entry.Value.RegisteredBy == "" {
			contactInfoString = "this link was registeredBy blank user. please specify empty string for delete this entry"
		}
		status := http.StatusBadRequest
		http.Error(w, http.StatusText(status)+". "+key+" is not registeredBy "+registeredBy+". "+contactInfoString, status)
		return
	}

	err = operations.RunRm(filepath, key, registeredBy)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	return
}
