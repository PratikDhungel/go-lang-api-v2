package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URL string
	Method string
	Handler func(http.ResponseWriter, *http.Request)
}


func LoadRoutes() []Route {

	routes := userRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router{

	for _, route := range LoadRoutes(){

		r.HandleFunc(route.URL, route.Handler).Methods(route.Method)
	}
	return r
}
