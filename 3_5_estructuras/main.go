package main

import "fmt"

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func main() {
	alexys := Person{
		Name:        "Alexys",
		Age:         42,
		HasChildren: false,
	}

	// beto := Person{"Beto", 33, true}
	// alejo := Person{Age: 32}

	alvaro := &alexys
	alvaro.Age = 60

	// fmt.Printf("%+v\n", alexys.Name)
	// fmt.Printf("%+v\n", beto.HasChildren)
	// fmt.Printf("%+v\n", alejo.Age)
	fmt.Printf("%+v\n", alexys)
	fmt.Printf("%+v\n", alvaro)
}
