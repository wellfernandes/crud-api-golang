package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route represents all api routes
type Route struct {
	URI     string
	Method  string
	Request func(http.ResponseWriter, *http.Request)
	Auth    bool
}

// ConfigRoutes insert all routes into the router
func ConfigRoutes(router *mux.Router) *mux.Router {

	routes := UserRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Request).Methods(route.Method)
	}
	return router
}
