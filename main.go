package main

import (
	"github.com/DarManuals/clean-arch/handlers"
	"github.com/DarManuals/clean-arch/logger"
	"github.com/DarManuals/clean-arch/repository"
	"github.com/DarManuals/clean-arch/services"
)

const ServiceName = `UserService`

func main() {
	var (
		log          = logger.New(ServiceName)
		userRepo     = repository.NewUserStorage()
		userService  = services.NewUserService(userRepo)
		usersHandler = handlers.NewUserHandler(userService, log)
	)

	NewServer(":12345", NewRouter(usersHandler)).Run()
}
