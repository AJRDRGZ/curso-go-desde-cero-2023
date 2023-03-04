package main

import (
	"fmt"
	"strings"
)

func main() {
	result := sum(2, 3)
	fmt.Println(result)

	lower, upper := convert("EDteam")
	fmt.Println(lower, upper)
}

func convert(text string) (lower string, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)

	return
}

func sum(a, b int) int {
	return a + b
}
