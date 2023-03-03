package main

import "fmt"

func main() {
	character := "🧞"
	canSearch := true

	switch {
	case !canSearch:
		fmt.Println("no esta permitida la busqueda")
	case character == "🦸" || character == "🧞":
		fmt.Println("es un superheroe")
	case character == "🦹" || character == "🧟":
		fmt.Println("es un supervillano")
	default:
		fmt.Println("es un personaje normal")
	}
}
