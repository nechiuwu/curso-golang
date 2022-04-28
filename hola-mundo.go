package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	holaGo()
	//Toma valores que paso por defecto al ejecutar app
	fmt.Println("Hola " + os.Args[1] + "!")
	convertValue, err := strconv.Atoi(os.Args[2])
	if err == nil {
		fmt.Print("Valor convertido a int: ")
		fmt.Println(convertValue)
	} else {
		fmt.Println(err)
	}
	// A medida que se ejecuta ingreso datos
	fmt.Print("¿Cuál es tu edad?")
	var edad int
	fmt.Scanln(&edad)
	if edad >= 18 {
		fmt.Print("Eres MAYOR de edad")
	} else {
		fmt.Print("Eres MENOR de edad")
	}
}

func holaGo() {
	fmt.Println("****Mi programa con GO!****")
}
