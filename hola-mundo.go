package main

import (
	"fmt"
)

func main() {
	holaGo()
	var peliculas [3][2]string
	peliculas[0][0] = "El Retorno del Rey"
	peliculas[0][1] = "Aladdin"
	peliculas[1][0] = "Nemo"
	peliculas[1][1] = "El Rey Escorpión"
	peliculas[2][0] = "Paprika"
	peliculas[2][1] = "Piola"

	fmt.Print("Todas las pelis: ")
	fmt.Println(peliculas)
}

func holaGo() {
	fmt.Println("Hola desde GO!")
}
