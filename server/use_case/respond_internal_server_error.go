package use_case

import "net/http"

func respondInternalServerError(w http.ResponseWriter) {
	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)
	return
}
