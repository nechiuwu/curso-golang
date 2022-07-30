package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var moviesGlobal = Movies{
	Movie{"Titanic", 1997, "Spielberg"},
	Movie{"Jurassic Park", 1999, "Spielberg"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "devuelve texto")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, contactame")
}

func ListMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listado de peliculas")
	movies := moviesGlobal
	json.NewEncoder(w).Encode(movies)
}

func ShowMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]
	fmt.Fprintf(w, "Has cargado la pelicula %s", movieId)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movieData Movie
	// cuando usa & la variable aun no tiene nada y es la que tiene que rellener
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movieData)
	// a la primera variable le añado el segundo parametro de entrada
	moviesGlobal = append(moviesGlobal, movieData)
}
