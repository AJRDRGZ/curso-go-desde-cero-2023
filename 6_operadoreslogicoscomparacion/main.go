package main

import "fmt"

func main() {
	// Operadores Comparación: >, <, ==, !=, >=, <=
	fmt.Println(6 >= 4)

	// Operadores Lógicos &&, ||
	var age uint = 33
	fmt.Println("Es Adulto?:", age >= 18 && age <= 60)
	fmt.Println("Es Niño o Anciano?:", age < 18 || age > 60)

	// Operador lógico Unario: !
	fmt.Println(!(4 != 4))
}
