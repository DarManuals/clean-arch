package main

import (
	"github.com/darmanuals/clean-arch/DAL/mocked"
	userH "github.com/darmanuals/clean-arch/handlers/users"
	userS "github.com/darmanuals/clean-arch/services/users"
	"net/http"
)

func main() {
	var userDAO = mocked.NewUserStorage()
	var userService = userS.NewUserService(userDAO)
	var usersHandler = userH.NewUserHandler(userService)

	router := NewRouter(usersHandler)
	http.ListenAndServe(":12345", router)
}
