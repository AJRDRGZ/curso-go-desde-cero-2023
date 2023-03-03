package main

import "fmt"

func main() {
	// for clasico
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// for (while)
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// for forever
	i = 1
	for i <= 5 {
		if i == 6 {
			break
		}

		fmt.Println(i)
		i++
	}

	// for range slice
	for i, v := range []string{"🍕", "🍔", "🍎", "🌭"} {
		fmt.Println("indice:", i, "value:", v)
	}

	// for range slice (actualizar valores)
	numbers := []uint8{2, 4, 6, 8}
	for i := range numbers {
		numbers[i] *= 2
	}
	fmt.Println(numbers)

	// for maps
	food := map[string]string{
		"pizza":     "🍕",
		"hamburger": "🍔",
		"apple":     "🍎",
		"hotdog":    "🌭",
	}
	for key, value := range food {
		fmt.Println("key:", key, "value:", value)
	}

	// for range strings
	for i, v := range "EDteam" {
		fmt.Println("indice:", i, "value:", string(v))
	}
}
