/*desarrole una aplicacion que pida in dato y lo muestrteen la consola
 */
package main

import (
	"fmt"
)

func main() {
	var nom string
	//pedimos al usuario que ingrese su nombre
	fmt.Print("Ingrese el nombre: ")
	fmt.Scanln(&nom)
	//pdimos al usuario que ingrese su edad
	fmt.Print("Ingrese su edad: ")
	var ed int
	fmt.Scan(&ed)

	fmt.Println("su nombre es: ", nombre(nom))
	fmt.Println("Su edad es: ", edad(ed))
}
func nombre(nombre string) string {
	return nombre
}
func edad(edad int) int {
	return edad
}
