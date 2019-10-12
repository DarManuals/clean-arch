package services

import "github.com/DarManuals/clean-arch/models"

type User interface {
	Retrieve(id int) (models.User, error)
}
