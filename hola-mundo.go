package main

import (
	"fmt"
)

func main() {
	holaGo()
	fmt.Print("Pedido 1 --> ")
	fmt.Println(gorras(8, "CLP"))
	fmt.Print("Pedido 2 --> ")
	fmt.Println(gorras(5, "CLP"))
}

func gorras(pedido float32, moneda string) (string, float32, string) {
	precio := func() float32 {
		return pedido * 7
	}
	// precio pasa a ser una funcion
	return "El precio de gorras pedidas es:", precio(), moneda
}

func holaGo() {
	fmt.Println("Hola desde GO!")
}
