package main

import (
	"example/go-nurse-mgmt/pkg/database"
	"example/go-nurse-mgmt/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func setupRouter() http.Handler {
	r := mux.NewRouter()
	http.Handle("/", r)

	routes.NurseHandlerRoutes(r)
	routes.AuthRoutes(r)
	// CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	return corsMiddleware(r)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := setupRouter()

	//Calling for database connection
	database.DataMigration()

	//Creating Server on 8080
	http.ListenAndServe(":8080", router)
}
