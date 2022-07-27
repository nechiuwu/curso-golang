package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// crea endpoint y que hará cuando sea invocado
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "devuelve texto")
	})

	// crea un server
	server := http.ListenAndServe(":8080", nil)
	log.Fatal(server)

	fmt.Println("El server está corriendo en el puerto :8080")
}
