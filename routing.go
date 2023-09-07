package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func HandlerRouting(){
	r:= mux.NewRouter()
	r.HandleFunc("/nurses",GetNurse).Methods("GET")
	r.HandleFunc("/nurse/{id}",GetNurseById).Methods("GET")
	r.HandleFunc("/nurse",CreateNurse).Methods("POST")
	r.HandleFunc("/nurse/{id}",UpdateNurse).Methods("PUT")
	r.HandleFunc("/nurse/{id}",DeleteNurse).Methods("DELETE")
	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)
}