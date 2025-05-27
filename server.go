package main

import (
	"net/http"
)

type Server struct {
	Handler http.ServeMux
	Addr    string
}

// New Server -
func NewServer() Server {
	return Server{
		Handler: *http.NewServeMux(),
		Addr:    ":8080",
	}
}
