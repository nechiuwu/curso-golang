package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("-----Lector-----")
	writeFile()
	readFile()
}

func showError(e error) {
	if e != nil {
		panic(errors.New(e.Error()))
	}
}

func writeFile() {
	nuevoTexto := os.Args[1]
	// escribe := ioutil.WriteFile("lenguajes.txt", nuevoTexto, 0755)
	fichero, err := os.OpenFile("lenguajes.txt", os.O_APPEND, 0777)
	showError(err)

	escribe, err := fichero.WriteString(string("\n" + "5. " + nuevoTexto))
	fichero.Close()
	showError(err)
	fmt.Println(escribe)
}

func readFile() {
	lng, err := ioutil.ReadFile("lenguajes.txt")
	fmt.Println("Lenguajes de programacion", string(lng))
	showError(err)
}
