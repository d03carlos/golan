package main

import (
	"fmt"
)

func calcular(a, b float64, op string) float64 {
	var resultado float64
	switch op {
	case "+":
		resultado = a + b
	case "-":
		resultado = a - b
	case "*":
		resultado = a * b
	case "/":
		resultado = a / b
	default:
		fmt.Println("Operación no válida")
	}
	return resultado
}

func main() {
	var a, b float64
	var op string
	var resultado float64
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scan(&a)

	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scan(&b)

	fmt.Println("Ingrese el tipo de operación: suma: + resta: - multpilicacion: *  division: /  Raizcuadrada:v  ")
	fmt.Scan(&op)
	resultado = calcular(a, b, op)

	fmt.Println("El resultado de: ", a, op, b, " = ", resultado)
}
