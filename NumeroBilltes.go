package main

import "fmt"

func main() {
	var numBilletes int
	var numBilletes100 int
	var numBilletes50 int
	var numBilletes20 int
	var numBilletes10 int
	var numBilletes5 int
	var numBilletes2 int
	var numBilletes1 int
	var residuo int
	//Pedimos al usuario que ingrese el numero de billetes
	fmt.Println("Ingrese el numero de billetes")
	fmt.Scan(&numBilletes)
	//Calculamos el numero de billetes de cada denominacion
	numBilletes100 = numBilletes / 100
	residuo = numBilletes % 100
	numBilletes50 = residuo / 50
	residuo = residuo % 50
	numBilletes20 = residuo / 20
	residuo = residuo % 20
	numBilletes10 = residuo / 10
	residuo = residuo % 10
	numBilletes5 = residuo / 5
	residuo = residuo % 5
	numBilletes2 = residuo / 2
	residuo = residuo % 2
	numBilletes1 = residuo / 1
	//Mostramos el resultado
	fmt.Println("El numero de billetes de 100 es: ", numBilletes100)
	fmt.Println("El numero de billetes de 50 es: ", numBilletes50)
	fmt.Println("El numero de billetes de 20 es: ", numBilletes20)
	fmt.Println("El numero de billetes de 10 es: ", numBilletes10)
	fmt.Println("El numero de billetes de 5 es: ", numBilletes5)
	fmt.Println("El numero de billetes de 2 es: ", numBilletes2)
	fmt.Println("El numero de billetes de 1 es: ", numBilletes1)
}
