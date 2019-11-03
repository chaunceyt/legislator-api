package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// NewRouter responsibe for endpoints.
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

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
		"Metrics",
		"GET",
		"/metrics",
		Metrics,
	},
	Route{
		"LegislatorsIndex",
		"GET",
		"/",
		Index,
	},
	Route{
		"LegislatorsSearch",
		"POST",
		"/find-legislator",
		LegislatorsSearch,
	},
	Route{
		"LegislatorsIndex",
		"GET",
		"/legislators",
		LegislatorsIndex,
	},
	Route{
		"LegislatorsInfo",
		"GET",
		"/legislator/{firstname}/{lastname}/offices",
		LegislatorsInfo,
	},
	Route{
		"LegislatorsJSON",
		"GET",
		"/legislator/{firstname}/{lastname}/json",
		LegislatorsJSON,
	},
	Route{
		"LegislatorsHome",
		"GET",
		"/legislator/{firstname}/{lastname}",
		LegislatorsHome,
	},
}
