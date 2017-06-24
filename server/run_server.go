package server

import (
	"encoding/json"
	"fmt"
	"github.com/matsu-chara/gol/operations"
	"net/http"
	"strings"
)

// RunServer run server
func RunServer(filepath string, port uint) error {
	handler := makeHandler(filepath)
	http.HandleFunc("/", handler)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("starting gol server at %s.\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}

func makeHandler(filepath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(r.URL.Path, "/")
		key := params[1]

		respondGet(filepath, key, params[2:], w, r)
		return
	}
}

func respondGet(filepath string, key string, params []string, w http.ResponseWriter, r *http.Request) {
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

func respondInternalServerError(w http.ResponseWriter) {
	status := http.StatusInternalServerError
	http.Error(w, http.StatusText(status), status)
	return
}
