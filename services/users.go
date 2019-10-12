package services

import (
	"github.com/DarManuals/clean-arch/models"
	"github.com/DarManuals/clean-arch/repository"
)

type UserService struct {
	users repository.User
}

func NewUserService(userDAO repository.User) User {
	return UserService{
		users: userDAO,
	}
}

func (s UserService) Retrieve(id int) (models.User, error) {
	return s.users.Get(id)
}
