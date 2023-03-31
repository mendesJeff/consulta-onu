package router

import (
	"database/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate retorna um router com as configurações
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}
