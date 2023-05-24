package main

import "fmt"

func main() {
	//declaramos nuestras variables
	var base float64
	var altura float64
	//pedimos al usuario que ingrese la base
	fmt.Println("Ingrese la base:")
	fmt.Scan(&base)
	//pedimos al usuario que ingrese la altura
	fmt.Println("Ingrese la altura:")
	fmt.Scan(&altura)
	fmt.Println("El ára del rectangulo es: ", AreaTriangulo(base, altura))
}
func AreaTriangulo(b, h float64) float64 {
	return (b * h) / 2
}
