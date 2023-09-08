package routes

import (
	"github.com/gorilla/mux"
)

func NurseHandlerRouting(r *mux.Router){
	r.HandleFunc("/nurses",GetNurse).Methods("GET")
	r.HandleFunc("/nurse/{id}",GetNurseById).Methods("GET")
	r.HandleFunc("/nurse",CreateNurse).Methods("POST")
	r.HandleFunc("/nurse/{id}",UpdateNurse).Methods("PUT")
	r.HandleFunc("/nurse/{id}",DeleteNurse).Methods("DELETE")
	
}