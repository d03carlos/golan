package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("suma: ", suma(10, 2))
	fmt.Println("resta", resta(10, 2))
	fmt.Println("multiplicacion", multiplicacion(15, 2))
	resultado, err := division(15, 0)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("El resultado de la división es:", resultado)
	}
}
func suma(a, b int) int {
	return a + b
}
func resta(a, b int) int {
	return a - b
}
func multiplicacion(a, b int) int {
	return a * b
}
func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir entre 0")
	}
	return a / b, nil
}
