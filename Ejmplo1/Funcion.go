package main

import "fmt"

func suma(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("La suma es: ", suma(10, 20))
	fmt.Println(resta(8, 6))
	fmt.Println(division(10, 20))
	fmt.Print(multiplicacion(5, 6))
}
func resta(num1, num2 int) int {
	return num1 - num2
}
func multiplicacion(num1, num2 int) int {
	return num1 * num2
}
func division(num1, num2 int) int {
	return num1 / num2
}
