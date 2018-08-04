package main

import (
	"github.com/darmanuals/clean-arch/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(
	userHandler handlers.User,
) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.Get)
	return r
}
