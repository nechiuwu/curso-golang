package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		r.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			HandlerFunc(route.HandleFunc)
	}
	return r
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"Contact", "GET", "/contacto", Contact},
	Route{"ListMovies", "GET", "/peliculas", ListMovies},
	Route{"ShowMovie", "GET", "/peliculas/{id}", ShowMovie},
}
