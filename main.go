package main

import (
	"log"
	"net/http"
)

func main() {
	// URLs amigables usando /
	router := NewRouter()

	// Levanta conexion a DB
	connectMongoDb()

	// crea un server y recibe router
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
