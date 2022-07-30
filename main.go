package main

import (
	"log"
	"net/http"
)

func main() {
	// URLs amigables usando /
	router := NewRouter()

	// crea un server y recibe router
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
