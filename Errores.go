package main

import "fmt"

func main() {

	//solicitamos al usuario que ingrese dos numeros
	var num1, num2 int
	fmt.Println("Ingrese un numero")
	fmt.Scanln(&num1)
	fmt.Println("Ingrese otro numero")
	fmt.Scanln(&num2)
	//llamr a la fucnión  divide y le pasamos los dos numeros
	resultado, err := divide(float64(num1), float64(num2))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}
