package main

import "fmt"

func main() {
	if character := "🦹"; character == "🦸" {
		fmt.Println("es un superheroe")
	} else if character == "🦹" {
		fmt.Println("es un supervillano")
	} else {
		fmt.Println("es un personaje normal")
	}
}
