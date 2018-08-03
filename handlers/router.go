package handlers

import (
	"fmt"
	"github.com/darmanuals/clean-arch/services"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(
	users services.UserService,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, users.Retrieve(0))
	})

	return r
}
