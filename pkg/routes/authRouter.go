package routes

import (
	"example/go-nurse-mgmt/pkg/controller"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/signup", controller.SignUp).Methods("POST")
}
