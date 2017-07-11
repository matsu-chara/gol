package server

import (
	"fmt"
	"log"
	"net/http"
)

// RunServer run server
func RunServer(filepath string, port uint) error {
	handler := NewGolServerHandler(filepath)
	http.HandleFunc("/", handler)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("starting gol server at %s.\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	return nil
}
