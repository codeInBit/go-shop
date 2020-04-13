package controllers

import (
	"github.com/codeinbit/go-shop/api/middlewares"
)

func (s *Server) LoadRoutes() {
	//Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Admin routes
	s.Router.HandleFunc("/admins", middlewares.SetMiddlewareJSON(s.CreateAdmin)).Methods("POST")
	s.Router.HandleFunc("/admins", middlewares.SetMiddlewareJSON(s.GetAllAdmins)).Methods("GET")

}
