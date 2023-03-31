package routes

import (
	"database/src/controllers"
	"net/http"
)

var routesDatabase = []Route{
	{
		URI:                "/onus",
		Method:             http.MethodGet,
		Func:               controllers.SearchOnu,
		NeedAuthentication: false,
	},

	{
		URI:                "/olts",
		Method:             http.MethodGet,
		Func:               controllers.SearchByOlt,
		NeedAuthentication: false,
	},
}
