package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "alexys"
	toUpperCase(&name)

	fmt.Println(name)
	// greet("Alexys", "Lozada")
}

func toUpperCase(text *string) {
	*text = strings.ToUpper(*text)
}

// func greet(firstName, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }
