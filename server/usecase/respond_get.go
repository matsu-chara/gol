package usecase

import (
	"net/http"
	"strings"

	"github.com/matsu-chara/gol/operations"
)

// Get redirect to value url
func Get(filepath string, key string, params []string, w http.ResponseWriter, r *http.Request) {
	entry, err := operations.RunGet(filepath, key)
	if err != nil {
		respondInternalServerError(err, w)
		return
	}
	if entry == nil {
		http.NotFound(w, r)
		return
	}

	newURL := strings.Join(append([]string{entry.Value.Link}, params...), "/")
	http.Redirect(w, r, newURL, http.StatusSeeOther)

	return
}
