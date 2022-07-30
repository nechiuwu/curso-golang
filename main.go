package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// URLs amigables usando /
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)
	router.HandleFunc("/peliculas", ListMovies)
	router.HandleFunc("/peliculas/{id}", ShowMovie)

	// crea un server y recibe router
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

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
