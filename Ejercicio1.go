/*
Desarrolla una aplicacion que que solicite al usuario 5 numeros enteros
y no imprima en la consola el número mayor, menor, y el promedio
*/
package main

import "fmt"

func main() {
	//declaramos nuestras variables locales
	var num1 int
	var num2 int
	var num3 int
	var num4 int
	var num5 int
	var mayor int
	var menor int
	var promedio int
	//pedimos al usuario que ingrese los valores
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&num1)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&num2)
	fmt.Print("Ingrese el tercer numero: ")
	fmt.Scan(&num3)
	fmt.Print("Ingrese el cuarto numero: ")
	fmt.Scan(&num4)
	fmt.Print("Ingrese el quinto numero: ")
	fmt.Scan(&num5)
	//calculamos el mayor

	mayor = num1
	if num2 > mayor {
		mayor = num2
	}
	if num3 > mayor {
		mayor = num3
	}
	if num4 > mayor {
		mayor = num4
	}
	if num5 > mayor {
		mayor = num5
	}
	//calculamos el menor

	menor = num1
	if num2 < menor {
		menor = num2
	}
	if num3 < menor {
		menor = num3
	}
	if num4 < menor {
		menor = num4
	}
	if num5 < menor {
		menor = num5
	}
	//calculamos el promedio
	promedio = (num1 + num2 + num3 + num4 + num5) / 5
	//mostramos los resultados
	fmt.Println("El numero mayor es: ", mayor)
	fmt.Println("El numero menor es: ", menor)
	fmt.Println("El promedio es: ", promedio)

}
