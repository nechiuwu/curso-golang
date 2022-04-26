package main

import (
	"fmt"
)

type Gorra struct {
	marca  string
	color  string
	plana  bool
	precio float32
}

func main() {
	var jockey = Gorra{"nike", "verde", true, 20.5}
	fmt.Println(jockey)
}
