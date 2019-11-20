package mock

import "github.com/DarManuals/clean-arch/models"

type UserRepository struct {
	GetFn func(id int) (models.User, error)
}

func (u UserRepository) Get(id int) (models.User, error) {
	return u.GetFn(id)
}
