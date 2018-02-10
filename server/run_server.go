package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matsu-chara/gol/kvs"
)

// RunServer run server
func RunServer(filepath string, port uint) error {
	kvs.CacheWarming(filepath)
	handler := NewGolServerHandler(filepath)
	http.HandleFunc("/", handler)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("starting gol server at %s.\n", addr)
	return http.ListenAndServe(addr, nil)
}
