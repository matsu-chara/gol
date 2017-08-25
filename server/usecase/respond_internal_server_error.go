package usecase

import (
	"log"
	"net/http"
)

func respondInternalServerError(err error, w http.ResponseWriter) {
	log.Println(err)

	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)
	return
}
