package main

import "fmt"

func main() {
	//declaramos nuestras variables
	var edad int
	var nombre string

	fmt.Println("Ingrese el edad: ")
	fmt.Scan(&edad)

	fmt.Println("Ingrese el nombre: ")
	fmt.Scan(&nombre)

	fmt.Println("Su nombre es: ", nombre, " y tienes: ", edad, " años")
}
