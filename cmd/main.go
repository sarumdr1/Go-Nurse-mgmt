package main

import (
	"example/go-nurse-mgmt/pkg/database"
	"example/go-nurse-mgmt/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	//Calling for database connection
	database.DataMigration()

	//Caling Routes
	routes.NurseHandlerRoutes(r)
	routes.AuthRoutes(r)

	//Creating Server on 8080
	http.ListenAndServe(":8080", nil)
}
