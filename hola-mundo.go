package main

import (
	"fmt"
)

func main() {
	holaGo()
	peliculas := []string{
		"El Retorno del Rey",
		"Aladdin",
		"Nemo",
		"El Rey Escorpión",
		"Paprika",
		"Piola"}

	// Añade elementos al slice
	peliculas = append(peliculas, "Batman")

	fmt.Print("Todas las pelis: ")
	fmt.Println(peliculas)
}

func holaGo() {
	fmt.Println("Hola desde GO!")
}
