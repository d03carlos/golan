package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int
	num2 = num1
	num2 += num2
	fmt.Println(num1, num2)
}
