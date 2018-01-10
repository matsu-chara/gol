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
		} else if r.Method == "PUT" {
			lockCtx.Lock()
			defer lockCtx.Unlock()
			handlePut(filepath, w, r)
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
		usecase.DumpAsHTML(filepath, w)
		return
	}
	if params[1] == "api" && params[2] == "dump" {
		usecase.Dump(filepath, w)
		return
	}

	key := params[1]
	rest := params[2:]

	usecase.Get(filepath, key, rest, w, r)
	return
}

// put key and value from form url encoded parameter. format: value=foo
func handlePut(filepath string, w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	if len(params) != 2 {
		http.Error(w, "Error: key contains '/'. please specify key without '/'", http.StatusBadRequest)
		return
	}
	key := params[1]
	link := r.PostFormValue("link")
	if link == "" {
		link = r.PostFormValue("value") // for api compatibility from v0.3.0
	}
	if link == "" {
		http.Error(w, "Error: value is empty. please specify value in body and set Content-Type application/x-www-form-urlencoded", http.StatusBadRequest)
		return
	}
	registeredBy := r.PostFormValue("registeredBy")
	isForce := (r.PostFormValue("force") == "true")

	usecase.Put(filepath, key, link, registeredBy, isForce, w, r)
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

	registeredBy := r.URL.Query().Get("registeredBy")
	usecase.Delete(filepath, key, registeredBy, w, r)
	return
}

func handleUnsupported(filepath string, w http.ResponseWriter, r *http.Request) {
	status := http.StatusMethodNotAllowed
	http.Error(w, http.StatusText(status), status)
	return
}
