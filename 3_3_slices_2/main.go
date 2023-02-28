package main

import "fmt"

func main() {
	// len(): # de elementos en el slice
	// cap(): # de elementos del array origen, a partir del índice donde se creo el slice

	// animals := [5]string{"🦍", "🐕", "🦮", "🐦", "🐘"}
	// pets := animals[1:3] // "🐕", "🦮"
	// pets = append(pets, "🐈", "🐕", "🐱")

	// Array[4]{"🐕", "🦮", "🐦", "🐘"}
	// new Array[8]{"🐕", "🦮",  "🐈", "🐕", "🐱"}

	// fmt.Println("animals:", animals)
	// fmt.Println("pets:", pets)
	// fmt.Println("tamaño pets:", len(pets))
	// fmt.Println("capacidad pets:", cap(pets))

	// pets := []string{"🐕", "🦮"}
	// pets := make([]string, 0, 3)
	// pets = append(pets, "🐈", "🐕", "🐱", "🐕")
	pets := []string{}

	fmt.Println("pets:", pets)
	fmt.Println("tamaño pets:", len(pets))
	fmt.Println("capacidad pets:", cap(pets))
	fmt.Println("valor cero:", pets == nil)
}
