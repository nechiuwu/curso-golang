package main

import (
	"fmt"
)

func main() {
	holaGo()
	var numero1 float32 = 3.5
	var numero2 float32 = 4.2
	calculadora(numero1, numero2)
}

func holaGo() {
	fmt.Println("Hola desde GO!")
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32
	switch op {
	case "+":
		resultado = n1 + n2
		break
	case "-":
		resultado = n1 - n2
		break
	case "*":
		resultado = n1 * n2
		break
	case "/":
		resultado = n1 / n2
		break
	default:
		resultado = 0
		break
	}
	return resultado
}

func calculadora(numero1 float32, numero2 float32) {
	fmt.Print("La suma es: ")
	fmt.Println(operacion(numero1, numero2, "+"))

	fmt.Print("La resta es: ")
	fmt.Println(operacion(numero1, numero2, "-"))

	fmt.Print("La multiplicación es: ")
	fmt.Println(operacion(numero1, numero2, "*"))

	fmt.Print("La división es: ")
	fmt.Println(operacion(numero1, numero2, "/"))
}
