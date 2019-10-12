package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DarManuals/clean-arch/services"
)

type UserHandler struct {
	users services.User
}

func NewUserHandler(us services.User) User {
	return UserHandler{
		users: us,
	}
}

func (h UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(h.users.Retrieve(0)); err != nil {
		log.Printf("User.Get err: %v", err)
	}
}
