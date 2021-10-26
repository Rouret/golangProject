package api

import (
	mux "github.com/julienschmidt/httprouter" //using HTTPRouter package
)

type Route struct {
        Method string
        Path   string
        Handle mux.Handle   // httprouter package as mux
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/messages",
		MessageIndex,
	},
	Route{
		"GET",
		"/messages/:id",
		MessageShow,
	},
	Route{
		"POST",
		"/messages",
		MessageCreate,
	},
}

