package main

import (
	"github.com/DarManuals/clean-arch/handlers"
	"github.com/DarManuals/clean-arch/repository"
	"github.com/DarManuals/clean-arch/services"
)

func main() {
	var (
		userRepo     = repository.NewUserStorage()
		userService  = services.NewUserService(userRepo)
		usersHandler = handlers.NewUserHandler(userService)
	)

	NewServer(":12345", NewRouter(usersHandler)).Run()
}
