package main

import (
	"fmt"
)

func main() {
	var nombre string
	//Pedimos al usuario que ingrese su nombre
	fmt.Println("Ingrese su nombre")
	//fmt.Scan(&nombre)
	fmt.Scanln(&nombre)
	//imprimos el nombre en consola
	fmt.Println("Hola", nombre)
}
