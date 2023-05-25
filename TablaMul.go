package main

import "fmt"

func main() {
	//con el bucle for
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 12; j++ {
			fmt.Println(i, " X ", j, "=", i*j)
		}
		fmt.Println("====================")
	}
	//con el bucle while
	fmt.Println("====================")
	var i int = 5
	for i <= 12 {
		var j int = 0
		for j <= 12 {
			fmt.Println(i, " X ", j, "=", i*j)
			j++
		}
		i++
		fmt.Println("====================	")
	}
}
