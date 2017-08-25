package server

import (
	"github.com/matsu-chara/gol/server/usecase"
	"net/http"
	"strings"
)

// NewGolServerHandler creates new handler
func NewGolServerHandler(filepath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			lockCtx.RLock()
			defer lockCtx.RUnlock()
			handleGet(filepath, w, r)
		} else if r.Method == "POST" {
			lockCtx.Lock()
			defer lockCtx.Unlock()
			handlePost(filepath, w, r)
		} else if r.Method == "DELETE" {
			lockCtx.Lock()
			defer lockCtx.Unlock()
			handleDelete(filepath, w, r)
		} else {
			handleUnsupported(filepath, w, r)
		}
		return
	}
}

// get from url path or dump
// params: first part is key, rest will be appended to value
func handleGet(filepath string, w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) <= 1 || params[1] == "" {
		usecase.Dump(filepath, w)
		return
	}

	key := params[1]
	rest := params[2:]

	usecase.Get(filepath, key, rest, w, r)
	return
}

// put key and value from form url encoded parameter. format: value=foo
func handlePost(filepath string, w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) != 2 {
		http.Error(w, "Error: key contains '/'. please specify key without '/'", http.StatusBadRequest)
		return
	}
	key := params[1]
	value := r.PostFormValue("value")
	if value == "" {
		http.Error(w, "Error: value is empty. please specify value in POST body and set Content-Type application/x-www-form-urlencoded", http.StatusBadRequest)
		return
	}

	usecase.Post(filepath, key, value, w, r)
	return
}

// delete key
func handleDelete(filepath string, w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) != 2 {
		http.Error(w, "Error: key contains '/'. please specify key without '/'", http.StatusBadRequest)
		return
	}
	key := params[1]

	usecase.Delete(filepath, key, w, r)
	return
}

func handleUnsupported(filepath string, w http.ResponseWriter, r *http.Request) {
	status := http.StatusMethodNotAllowed
	http.Error(w, http.StatusText(status), status)
	return
}
