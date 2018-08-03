package mocked

import (
	"github.com/darmanuals/clean-arch/DAL"
	"github.com/darmanuals/clean-arch/models"
)

type UserDAOimpl struct{}

func NewUserStorage() DAL.UserDAO {
	return UserDAOimpl{}
}

func (UserDAOimpl) Get(id int) *models.User {
	return &models.User{
		ID:   id,
		Name: "Test",
	}
}
