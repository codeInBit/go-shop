package routes

import(
	"github.com/gorilla/mux"
	"github.com/codeinbit/go-shop/api/middlewares"
)

type Route struct {
	Router *mux.Router
}

func (r Route) LoadRouter()  {
	r.Router = mux.NewRouter()

	//Home Route
	r.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}
