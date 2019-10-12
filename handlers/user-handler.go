package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DarManuals/clean-arch/logger"
	"github.com/DarManuals/clean-arch/services"
)

type UserHandler struct {
	users services.User
	log   logger.Logger
}

func NewUserHandler(us services.User, l logger.Logger) User {
	return UserHandler{
		users: us,
		log:   l,
	}
}

func (h UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Get: start handling request.")

	user, err := h.users.Retrieve(0)
	if err != nil {
		h.log.Error("Get: request failed", err)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.log.Error("Get: request failed ", err)
	}

	h.log.Info("Get: request finished successfully.")
}
