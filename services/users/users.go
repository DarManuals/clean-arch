package Users

import (
	"github.com/darmanuals/clean-arch/DAL"
	"github.com/darmanuals/clean-arch/models"
	"github.com/darmanuals/clean-arch/services"
)

type UserService struct {
	users DAL.UserDAO
}

func NewUserService(userDAO DAL.UserDAO) services.User {
	return UserService{
		users: userDAO,
	}
}

func (s UserService) Retrieve(id int) *models.User {
	return s.users.Get(id)
}
