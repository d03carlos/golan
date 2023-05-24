package main

import "fmt"

func main() {
	//declaramos nuestras variables
	var num1 int
	var num2 int
	//pedimos al suaurio que ingrese el primer numero
	fmt.Println("Ingrese el primer numero")
	fmt.Scanln(&num1) //capturamos el valor ingrsado por teclado y lo alamcenamos en la varible num1
	fmt.Println("Ingrese el segundo numero")
	fmt.Scanln(&num2) //capturamos el valor ingrsado por y lo alamcenamos en la varible num2
	//realizamos las operacione solicitadas
	fmt.Println("La suma es: ", num1+num2)
	fmt.Println("La resta es: ", num1-num2)
	fmt.Println("La multiplicacion es: ", num1*num2)
	fmt.Println("La division es: ", num1/num2)
}
