package main

import "fmt"

func main() {
	var numero int
	//pedimos al usuario que ingrese un numero
	fmt.Println("Ingrese un numero")
	fmt.Scan(&numero)
	switch numero {
	case 1:
		fmt.Println("Es Lunes")
	case 2:
		fmt.Println("Es Martes")
	case 3:
		fmt.Println("Es Miercoles")
	case 4:
		fmt.Println("Es Jueves")
	case 5:
		fmt.Println("Es Viernes")
	case 6:
		fmt.Println("Es Sabado")
	case 7:
		fmt.Println("Es Domingo")
		break
	}

}
