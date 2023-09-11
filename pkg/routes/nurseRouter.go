package routes

import (
	"example/go-nurse-mgmt/pkg/controller"
	"example/go-nurse-mgmt/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NurseHandlerRoutes(r *mux.Router) {
	r.Handle("/nurses", middleware.AuthMiddleware(http.HandlerFunc(controller.GetNurse))).Methods("GET")
	r.Handle("/nurse/{id}", middleware.AuthMiddleware(http.HandlerFunc(controller.GetNurseById))).Methods("GET")
	r.Handle("/nurse", middleware.AuthMiddleware(http.HandlerFunc(controller.CreateNurse))).Methods("POST")
	r.Handle("/nurse/{id}", middleware.AuthMiddleware(http.HandlerFunc(controller.UpdateNurse))).Methods("PUT")
	r.Handle("/nurse/{id}", middleware.AuthMiddleware(http.HandlerFunc(controller.DeleteNurse))).Methods("DELETE")

}
