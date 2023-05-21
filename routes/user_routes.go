package routes

import (
	"crud-api-golang/internal/controller"
	"crud-api-golang/rest"
	"net/http"
)

var userRoutes = []rest.Route{
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
