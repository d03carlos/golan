package main

import "fmt"

func main() {
	//Pedimos al usuario que ingrese un numero
	var n int
	fmt.Print("Ingrese el tamaño de la matriz: ")
	fmt.Scan(&n)
	var numeros = make([]int, n)
	/*for i := 0; i < n; i++ {
		fmt.Print("Ingrese un numero: ")
		fmt.Scan(&numeros[i])
	}*/
	var i int = 0
	for i < n {
		fmt.Print("Ingrese un numero: ")
		fmt.Scan(&numeros[i])
		i++
	}
	for i := 0; i < n; i++ {
		fmt.Println(numeros[i])
	}
}
