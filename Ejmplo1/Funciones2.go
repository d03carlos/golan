package main

import (
	"fmt"
)

func main() {
	//declaramos nuestras variables
	var num1 int
	var num2 int
	//pedimos al usuario que ingrese el primer numero
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&num1)
	//pedimos al usuario que ingrese el segundo numero
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&num2)
	//mostramos los esultados en la consola
	fmt.Println("La suma de: ", num1, " + ", num2, " = ", sumar(num1, num2))
	fmt.Println("La resta de: ", num1, " + ", num2, " = ", restar(num1, num2))
	fmt.Println("La multiplicación de: ", num1, " * ", num2, " = ", multiplicar(num1, num2))
	fmt.Println("La division de: ", num1, " / ", num2, " = ", dividir(num1, num2))
}
func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}
func dividir(a, b int) int {
	return a / b
}
