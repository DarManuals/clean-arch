package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DarManuals/clean-arch/logger"
	"github.com/DarManuals/clean-arch/services"

	"github.com/gorilla/mux"
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

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.log.Error("Get: request failed", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{Msg: err.Error()})
		return
	}

	user, err := h.users.Retrieve(id)
	if err != nil {
		h.log.Error("Get: request failed", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error{Msg: err.Error()})
		return
	}

	_ = json.NewEncoder(w).Encode(user)
	h.log.Info("Get: request finished successfully.")
}
