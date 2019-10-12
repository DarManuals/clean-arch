package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	address string
	router  *mux.Router
}

func (s Server) Run() {
	log.Printf("Server starting on [%s]\n", s.address)
	log.Fatalf("Server failed to start: %v\n", http.ListenAndServe(s.address, s.router))
}

func NewServer(address string, router *mux.Router) Server {
	return Server{address: address, router: router}
}
