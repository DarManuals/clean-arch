package handlers

import "net/http"

type User interface {
	Get(w http.ResponseWriter, r *http.Request)
}
