package repository

import "github.com/DarManuals/clean-arch/models"

type User interface {
	Get(id int) (models.User, error)
}
