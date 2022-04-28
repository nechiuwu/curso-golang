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
	fmt.Print("Cantidad de pelis: ")
	fmt.Println(len(peliculas))
	fmt.Print("Pelis del 0 al 3: ")
	fmt.Println(peliculas[0:3])
}

func holaGo() {
	fmt.Println("Hola desde GO!")
}
