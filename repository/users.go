package repository

import (
	"errors"

	"github.com/DarManuals/clean-arch/models"
)

type user struct{}

func NewUserStorage() User {
	return user{}
}

func (user) Get(id int) (models.User, error) {
	if id == 0 {
		return models.User{}, errors.New("mock err")
	}

	return models.User{
		ID:   id,
		Name: "Test",
	}, nil
}
