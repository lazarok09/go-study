package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a common interface to all api routes
type Route struct {
	URI            string
	Method         string
	Handler        func(http.ResponseWriter, *http.Request)
	IsAuthRequired bool
}

// Setup all routes inside the mux router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
