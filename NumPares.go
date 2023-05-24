package main

import "fmt"

func main() {
	//Declramos nuestras varibales locales
	var numero int
	var suma int = 0
	//pedimos al usuario que ingrese un numero
	fmt.Print("Ingrese un numero: ")
	//leemos el numero
	fmt.Scan(&numero)
	//imprims solo los numeros pares
	for i := 0; i <= numero; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			suma += i
		}
	}
	fmt.Println("La suma de los numeros pares es: ", suma)
}
