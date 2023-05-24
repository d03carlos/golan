package main

import "fmt"

func main() {
	//declramos nuestras variables
	var num1, num2 int
	var operador string
	//pedimos al usuario que ingrese el primer numero
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&num1)
	//pedimos al usuario que ingrese el segundo numero
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&num2)
	//Pedimos al usuario que ingrese el operador
	fmt.Print("Ingrese el operador:\n suma: + \n resta: -\n multiplicación: *\n division: /")
	fmt.Scan(&operador)
	//realizamos las operaciones solicitadas
	if operador == "+" {
		fmt.Println("La suma de:", num1, " + ", num2, " = ", sumar(num1, num2))
	}
	if operador == "-" {
		fmt.Println("La resta de:", num1, " - ", num2, " = ", restar(num1, num2))
	}
	if operador == "*" {
		fmt.Println("La multiplicación de:", num1, " * ", num2, " = ", multiplicar(num1, num2))
	}
	if operador == "/" {
		fmt.Println("La division de:", num1, " / ", num2, " = ", dividir(num1, num2))
	}

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
