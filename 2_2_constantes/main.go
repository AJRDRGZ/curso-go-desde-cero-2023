package main

import "fmt"

const (
	os     = "linux"
	domain = "ed.team"
)

const (
	Jan = iota + 1
	Feb
	Mar
	Abr
	May
	Jun
)

func main() {
	fmt.Println(May)
}
