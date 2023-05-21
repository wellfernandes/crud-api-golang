package rest

import "github.com/gorilla/mux"

// NewServerRouter returns a configured server router
func NewServerRouter() *mux.Router {

	return mux.NewRouter()
}
