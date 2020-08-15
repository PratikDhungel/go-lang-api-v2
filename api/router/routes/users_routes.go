package routes

import (
	"github.com/quality-dashboard/api/controllers"
	"net/http"
)

var userRoutes  = []Route{

	{
		URL: "/users/getAll",
		Method: http.MethodGet,
		Handler: controllers.GetAllUsers,
	},

	{
		URL: "/user/login/{id}",
		Method: http.MethodPost,
		Handler: controllers.UserLogin,
	},
}