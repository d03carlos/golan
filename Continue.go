package main

import "fmt"

func main() {
	//declarmos nuestra variable
	var numero int
	//pedimos al usuario que ingrese un numero
	fmt.Print("Ingrese un numero: ")
	fmt.Scan(&numero)
	for i := 1; i <= numero; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}

}
