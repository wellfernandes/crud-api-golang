package rest

import "net/http"

// Route represents all api routes
type Route struct {
	URI     string
	Method  string
	Request func(http.ResponseWriter, *http.Request)
	Auth    bool
}
