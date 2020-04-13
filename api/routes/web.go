package routes

import (
	"github.com/codeinbit/go-shop/api/controllers"
	"github.com/codeinbit/go-shop/api/middlewares"
	"github.com/gorilla/mux"
)

func LoadRouter() *mux.Router {
	var route *mux.Router
	route = mux.NewRouter()
	//Home Route
	route.HandleFunc("/", middlewares.SetMiddlewareJSON(controllers.Server{}.Home)).Methods("GET")

	//Admin routes
	route.HandleFunc("/users", middlewares.SetMiddlewareJSON(controllers.Server{}.CreateAdmin)).Methods("POST")
	route.HandleFunc("/users", middlewares.SetMiddlewareJSON(controllers.Server{}.GetAllAdmins)).Methods("GET")


	return route
}
