package services

import "github.com/darmanuals/clean-arch/models"

type User interface {
	Retrieve(id int) *models.User
}
