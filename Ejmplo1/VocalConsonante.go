/*
escriba una aplicaciom que muestres si es una vocal o consonante
el en valor ingresado por teclado
*/
package main

import (
	"fmt"
)

func main() {
	var letra string
	//pedimoa al usuario que ingrese una letra
	fmt.Println("Ingrese una letra")
	fmt.Scanln(&letra)
	if letra == "a" || letra == "e" || letra == "i" || letra == "o" || letra == "u" {
		fmt.Println("La letra introducida es una Vocal")
	} else {
		fmt.Println("La letra introducida es una Consonante")
	}
}
