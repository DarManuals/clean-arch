package handlers

import (
	"encoding/json"
	"github.com/darmanuals/clean-arch/handlers"
	"github.com/darmanuals/clean-arch/services"
	"net/http"
)

type UserHandlerImpl struct {
	users services.User
}

func NewUserHandler(us services.User) handlers.User {
	return UserHandlerImpl{
		users: us,
	}
}

func (h UserHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.users.Retrieve(0))
}
