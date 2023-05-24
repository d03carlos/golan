package main

import "fmt"

func main() {
	//declaramos nuestras variables locales
	var num int
	//pedimos al usuario que ingrese un numero
	fmt.Print("Ingrese un numero: ")
	fmt.Scan(&num)
	//realizamos las operacines solicitadas
	if num%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}
}
