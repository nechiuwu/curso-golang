package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Lector")
	lng, err := ioutil.ReadFile("lenguajes.txt")
	fmt.Println("Lenguajes de programacion", lng)
	showError(err)
}

func showError(e error) {
	if e != nil {
		panic(errors.New(e.Error()))
	}
}
