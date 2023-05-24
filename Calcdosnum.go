package main

import (
	"fmt"
)

func main() {
	//declaramos nuestras varibles locales
	var num1 int
	var num2 int
	var suma int
	var resta, mulplicacion, division int
	//pedimos los valores
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&num2)
	//realizamos las operaciones
	suma = num1 + num2
	resta = num1 - num2
	mulplicacion = num1 * num2
	division = num1 / num2
	//mostramos los resultados
	fmt.Println("La suma es: ", suma)
	fmt.Println("La resta es: ", resta)
	fmt.Println("La multiplicacion es: ", mulplicacion)
	fmt.Println("La division es: ", division)
}
