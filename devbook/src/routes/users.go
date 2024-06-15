package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Handler:        controllers.CreateUser,
		IsAuthRequired: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Handler:        controllers.SearchUsers,
		IsAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodGet,
		Handler:        controllers.SearchUser,
		IsAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodPut,
		Handler:        controllers.UpdateUser,
		IsAuthRequired: false,
	},
	{
		URI:            "/users/{id}",
		Method:         http.MethodDelete,
		Handler:        controllers.DeleteUser,
		IsAuthRequired: false,
	},
}
