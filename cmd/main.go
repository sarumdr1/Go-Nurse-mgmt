package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    http.Handle("/",r)

	http.ListenAndServe(":8080",nil)
    routes.NurseHandlerRouting(r)

    DataMigration()

}