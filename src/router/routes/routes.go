package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Describe all styles routes
type Routes struct {
	Uri         string
	Method      string
	Func        func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}

func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}

	return r
}
