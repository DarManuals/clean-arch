package services

import "github.com/darmanuals/clean-arch/models"

type UserService interface {
	Retrieve(id int) *models.User
}
