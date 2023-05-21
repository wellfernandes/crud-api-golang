package rest

import (
	"crud-api-golang/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
)

// NewServerRouter returns a configured server router
func NewServerRouter() *mux.Router {

	r := mux.NewRouter()
	return ConfigRoutes(r)
}

var UserRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Request: controller.Create,
		Auth:    false,
	},
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Request: controller.GetList,
		Auth:    false,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodGet,
		Request: controller.GetByID,
		Auth:    false,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodPut,
		Request: controller.Update,
		Auth:    false,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodDelete,
		Request: controller.Delete,
		Auth:    false,
	},
}
