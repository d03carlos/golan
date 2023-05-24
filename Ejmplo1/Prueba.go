package main

import "fmt"

func main() {
	var num int
	fmt.Print("Ingresa un número del 1 al 12: ")
	fmt.Scan(&num)

	switch num {
	case 12, 1, 2:
		fmt.Println("Invierno")
	case 3, 4, 5:
		fmt.Println("Primavera")
	case 6, 7, 8:
		fmt.Println("Verano")
	case 9, 10, 11:
		fmt.Println("Otoño")
	default:
		fmt.Println("Número inválido")
	}
}
