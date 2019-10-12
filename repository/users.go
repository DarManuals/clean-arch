package repository

import (
	"github.com/DarManuals/clean-arch/models"
)

type user struct{}

func NewUserStorage() User {
	return user{}
}

func (user) Get(id int) models.User {
	return models.User{
		ID:   id,
		Name: "Test",
	}
}
