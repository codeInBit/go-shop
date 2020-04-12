package routes

import (
	"github.com/codeinbit/go-shop/api/controllers"
	"github.com/codeinbit/go-shop/api/middlewares"
	"github.com/gorilla/mux"
)

func LoadRouter() {
	route := mux.NewRouter()
	//Home Route
	route.HandleFunc("/", middlewares.SetMiddlewareJSON(controller)).Methods("GET")
}
