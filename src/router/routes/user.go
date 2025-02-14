package routes

import (
	"api/src/controller"
	"net/http"
)

var usersRoutes = []Routes{
	{
		Uri:         "/usuarios",
		Method:      http.MethodPost,
		Func:        controller.CreateUser,
		RequestAuth: false,
	},
	{
		Uri:         "/usuarios",
		Method:      http.MethodGet,
		Func:        controller.SearchUsers,
		RequestAuth: false,
	},
	{
		Uri:         "/usuarios/{userId}",
		Method:      http.MethodGet,
		Func:        controller.SearchUser,
		RequestAuth: false,
	},
	{
		Uri:         "/usuarios/{userId}",
		Method:      http.MethodPut,
		Func:        controller.UpdateUser,
		RequestAuth: false,
	},
	{
		Uri:         "/usuarios/{userId}",
		Method:      http.MethodDelete,
		Func:        controller.DeleteUser,
		RequestAuth: false,
	},
}
