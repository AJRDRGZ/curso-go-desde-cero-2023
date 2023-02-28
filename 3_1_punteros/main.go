package main

import "fmt"

func main() {
	// Puntero: Variables que almacenan la dirección en memoria de un valor
	var color string = "🟥"
	var pointerColor *string
	pointerColor = &color
	*pointerColor = "🟦"

	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", color, color, &color)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", pointerColor, pointerColor, *pointerColor)
}
