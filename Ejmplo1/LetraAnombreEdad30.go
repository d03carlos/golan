package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	if strings.HasPrefix(name, "A") && age >= 30 {
		fmt.Println("Tu nombre empieza con 'A' y tienes 30 años o más.")
	} else {
		fmt.Println("Lo siento, no cumples los requisitos.")
	}
}
