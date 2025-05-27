package main

import (
	"log"
	"net/http"
)

func main() {
	cServer := NewServer()

	err := http.ListenAndServe(cServer.Addr, &cServer.Handler)
	if err != nil {
		log.Printf("couldn't start server: %v", err)
	}
}
