package main

import (
	"fmt"
	"strconv"
)

func main() {
	var entero int = 15
	//var String string = "12"
	var cadena string
	cadena = strconv.Itoa(entero)

	fmt.Println(entero, cadena)
}
