package swagger

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	sh := http.StripPrefix("/ui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/ui/").Handler(sh)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var routes = Routes{
	Route{
		"AddSecret",
		"POST",
		"/secret",
		AddSecret,
	},

	Route{
		"GetSecret",
		"GET",
		"/secret/{hash}",
		GetSecret,
	},
}
