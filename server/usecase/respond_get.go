package usecase

import (
	"fmt"
	"github.com/matsu-chara/gol/operations"
	"net/http"
	"strings"
)

func Get(filepath string, key string, params []string, w http.ResponseWriter, r *http.Request) {
	entry, err := operations.RunGet(filepath, key)
	if err != nil {
		fmt.Println(err)
		respondInternalServerError(w)
		return
	}
	if entry == nil {
		http.NotFound(w, r)
		return
	}

	newUrl := strings.Join(append([]string{entry.Value}, params...), "/")
	http.Redirect(w, r, newUrl, http.StatusSeeOther)

	return
}
