package main

import (
	"github.com/darmanuals/clean-arch/DAL/mocked"
	"github.com/darmanuals/clean-arch/handlers"
	"github.com/darmanuals/clean-arch/services/Users"
	"net/http"
)

func main() {
	var userDAO = mocked.NewUserStorage()
	var userService = Users.NewUserService(userDAO)

	router := handlers.NewRouter(userService) // todo: ?

	http.ListenAndServe(":12345", router)
}
