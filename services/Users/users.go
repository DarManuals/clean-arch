package Users

import (
	"github.com/darmanuals/clean-arch/DAL"
	"github.com/darmanuals/clean-arch/models"
	"github.com/darmanuals/clean-arch/services"
)

type UserServiceImpl struct {
	users DAL.UserDAO
}

func NewUserService(userDAO DAL.UserDAO) services.UserService {
	return UserServiceImpl{
		users: userDAO,
	}
}

func (s UserServiceImpl) Retrieve(id int) *models.User {
	return s.users.Get(id)
}
