package main

import "fmt"

func main() {
	var numeros = [5]int{1, 2, 3, 4, 5}
	for i, numero := range numeros {
		fmt.Println(i, numero)
	}
}
