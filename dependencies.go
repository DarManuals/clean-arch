package main

import (
	"github.com/DarManuals/clean-arch/handlers"
	"github.com/DarManuals/clean-arch/logger"
	"github.com/DarManuals/clean-arch/repository"
	"github.com/DarManuals/clean-arch/services"

	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

const (
	ServiceName = `UserService`
	defaultPort = `:12345`
)

func Initialize() *dig.Container {
	deps := dig.New()

	for _, constructor := range []interface{}{
		defaultLogger,
		defaultServer,
		NewRouter,
		repository.NewUserStorage,
		services.NewUserService,
		handlers.NewUserHandler,
	} {
		if err := deps.Provide(constructor); err != nil {
			panic(err)
		}
	}

	return deps
}

func defaultLogger() logger.Logger {
	return logger.New(ServiceName)
}

func defaultServer(r *mux.Router) Server {
	return NewServer(defaultPort, r)
}
