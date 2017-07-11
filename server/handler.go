package server

import (
	"github.com/matsu-chara/gol/server/use_case"
	"net/http"
	"strings"
)

func NewGolServerHandler(filepath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := strings.Split(r.URL.Path, "/")
		key := params[1]

		if key == "" {
			use_case.Dump(filepath, w)
			return
		}

		use_case.Get(filepath, key, params[2:], w, r)
		return
	}
}
