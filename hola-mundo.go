package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("El número: ", i, " es par")
		} else {
			fmt.Println("El número: ", i, "es impar")
		}
	}

	peliculas := []string{"Gran Torino", "The Book's thief", "Romeo + Juliet"}
	for i := 0; i < len(peliculas); i++ {
		if peliculas[i] == "The Book's thief" {
			fmt.Println("De seguro lloraste con", peliculas[i])
		} else {
			fmt.Println("Wenas pelis t gustan krnal", peliculas[i])
		}
	}

	for _, pelicula := range peliculas {
		fmt.Println("Recorriendo slice:", pelicula)
	}
}
