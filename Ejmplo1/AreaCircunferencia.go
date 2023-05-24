package main

import (
	"fmt"
	"math"
)

func main() {
	//declaramos nuestras variables
	var radio float64
	//pedimos al usuario que ingrese el radio
	fmt.Println("Ingrese el radio: ")
	fmt.Scan(&radio)
	fmt.Println("El área del curculo es: ", AreaCircunferencia(radio))
}
func AreaCircunferencia(ra float64) float64 {
	return math.Pi * ra * ra
}
