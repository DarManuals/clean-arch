package DAL

import "github.com/darmanuals/clean-arch/models"

type UserDAO interface {
	Get(int) *models.User
}
