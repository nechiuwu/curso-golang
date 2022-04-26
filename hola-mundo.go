package main

import (
	"fmt"
)

func main() {
	holaGo()

	/* Forma básica de definir array
	var peliculas [3]string
	peliculas[0] = "Moulin Rouge"
	peliculas[1] = "Pulp Fiction"
	peliculas[2] = "Gran Torino"
	fmt.Print("Todas las pelis: ")
	fmt.Println(peliculas)
	fmt.Print("La primera peli: ")
	fmt.Println(peliculas[0])*/

	peliculas := [3]string{
		"Mi amigo Alexis",
		"Romeo + Julieta",
		"Fear and Loathing in the Vegas"}
	fmt.Print("Todas las pelis: ")
	fmt.Println(peliculas)
}

func holaGo() {
	fmt.Println("Hola desde GO!")
}
