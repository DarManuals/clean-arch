package main

import (
	"github.com/DarManuals/clean-arch/config"
	"github.com/DarManuals/clean-arch/handlers"
	"github.com/DarManuals/clean-arch/logger"
	"github.com/DarManuals/clean-arch/repository"
	"github.com/DarManuals/clean-arch/services"

	"go.uber.org/dig"
)

func Initialize() *dig.Container {
	deps := dig.New()

	for _, constructor := range []interface{}{
		NewServer,
		NewRouter,
		config.New,
		logger.New,
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
