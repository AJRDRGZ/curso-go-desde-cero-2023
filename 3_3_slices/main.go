package main

import "fmt"

func main() {
	// Slice: son apuntadores a un Array, no poseen datos.
	things := [7]string{"🚓", "🚙", "🚕", "🏎", "🚒", "🚨", "🎈"}
	cars := things[:5] //  "🚓", "🚙", "🚕", "🏎", "🚒"
	red := things[3:]  // "🏎", "🚒", "🚨", "🎈"
	all := things[:]
	red[1] = "🚑"

	fmt.Println("Things:", things)
	fmt.Println("Cars:", cars)
	fmt.Println("Red:", red)
	fmt.Println("all:", all)
}
