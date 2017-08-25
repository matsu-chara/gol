package server

import (
	"github.com/matsu-chara/gol/server/usecase"
	"net/http"
	"strings"
)

// NewGolServerHandler creates new handler
func NewGolServerHandler(filepath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lockCtx.RLock()
		defer lockCtx.RUnlock()

		params := strings.Split(r.URL.Path, "/")
		key := params[1]

		if key == "" {
			usecase.Dump(filepath, w)
			return
		}

		usecase.Get(filepath, key, params[2:], w, r)
		return
	}
}
