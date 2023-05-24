package main

import (
	"fmt"
)

func main() {
	//declaramos nuestras variables
	var lado int
	fmt.Println("Ingrese la lado: ")
	fmt.Scan(&lado)
	fmt.Println("El área del cuadrado es: ", AreaCuadrado(lado))
}
func AreaCuadrado(l int) int {
	return l * l
}
