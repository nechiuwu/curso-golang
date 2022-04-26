package main

import (
	"fmt"
)

func main() {
	holaGo()
	pantalon("rojo", "largo", "jeans")
}

func pantalon(attrs ...string) {
	for _, attr := range attrs {
		fmt.Println(attr)
	}
}

func holaGo() {
	fmt.Println("Hola desde GO!")
}
