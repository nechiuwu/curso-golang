package main

import (
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

	/*
		// crea endpoint y que hará cuando sea invocado
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "devuelve texto")
		})
	*/

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
