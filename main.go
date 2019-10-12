package main

import (
	"net/http"

	"github.com/DarManuals/clean-arch/handlers"
	"github.com/DarManuals/clean-arch/repository"
	"github.com/DarManuals/clean-arch/services"
)

func main() {
	var userRepo = repository.NewUserStorage()
	var userService = services.NewUserService(userRepo)
	var usersHandler = handlers.NewUserHandler(userService)

	router := NewRouter(usersHandler)
	_ = http.ListenAndServe(":12345", router)
}
