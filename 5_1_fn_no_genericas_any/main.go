package main

import "fmt"

func main() {
	PrintList("EDteam", "gophers", "Go desde cero")
	PrintList(1, 2, 8)
}

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
