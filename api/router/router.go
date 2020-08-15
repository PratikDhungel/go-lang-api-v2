package router

import (
	"github.com/gorilla/mux"
	"github.com/quality-dashboard/api/router/routes"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
