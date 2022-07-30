package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "devuelve texto")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, contactame")
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listado de peliculas")
	movies := Movies{
		Movie{"Titanic", 1997, "Spielberg"},
		Movie{"Jurassic Park", 1999, "Spielberg"},
	}
	json.NewEncoder(w).Encode(movies)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w, "Has cargado la pelicula %s", movieId)
}
