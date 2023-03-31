package routes

import (
	"database/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// List of all Routes from API
type Route struct {
	URI                string
	Method             string
	Func               func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

// Config insert all routers inside the Router
func Configure(r *mux.Router) *mux.Router {
	routes := routesDatabase

	for _, route := range routes {

		if route.NeedAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
