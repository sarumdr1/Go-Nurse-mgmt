package routes

import(
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router){
r.HandleFunc("/login",Login).Methods("POST")
r.HandleFunc("/sign-up",SignUp).Methods("POST")
}