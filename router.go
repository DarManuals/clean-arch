package main

import (
	"github.com/gorilla/mux"

	"github.com/DarManuals/clean-arch/handlers"
)

func NewRouter(
	userHandler handlers.User,
) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.Get)
	return r
}
