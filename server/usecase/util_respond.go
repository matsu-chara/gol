package usecase

import (
	"github.com/matsu-chara/gol/kvs"
	"net/http"
)

func respondAsNotRegisteredBy(requestedRegisteredBy string, entry *kvs.Entry, w http.ResponseWriter) {
	contactInfoString := "please contact to " + entry.Value.RegisteredBy
	if entry.Value.IsRegisteredBy("") {
		contactInfoString = "this link was registeredBy blank user. please specify empty string for delete this entry"
	}
	status := http.StatusBadRequest
	http.Error(w, http.StatusText(status)+". "+entry.Key+" is not registeredBy "+requestedRegisteredBy+". "+contactInfoString, status)
	return
}
