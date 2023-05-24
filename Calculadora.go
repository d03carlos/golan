package main

import "fmt"

func main() {
	//Declaramos nuestras variables locales
	var num1, num2 int
	var resultado int
	var operador string
	//Pedimos los datos al usuario
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&num2)
	fmt.Println("Ingrese el operador: ")
	fmt.Scan(&operador)
	if operador == "+" {
		resultado = num1 + num2
	} else if operador == "-" {
		resultado = num1 - num2
	} else if operador == "*" {
		resultado = num1 * num2
	} else if operador == "/" {
		resultado = num1 / num2
	}
	fmt.Println("El resultado de: ", num1, "  ", operador, " ", num2, "es: ", resultado)
}
