package main

import (
	"fmt"
	"time"
)

func main() {
	momento := time.Now().Local()
	hoy := momento.Weekday()

	switch hoy {
	case 0:
		fmt.Println("Hoy es Lunes")
	case 1:
		fmt.Println("Hoy es Martes")
	case 2:
		fmt.Println("Hoy es Miércoles")
	case 3:
		fmt.Println("Hoy es Jueves")
	case 4:
		fmt.Println("Hoy es Viernes")
	default:
		fmt.Println("Hoy es otro día")
	}
}
