package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MyInt int
type MyIntV2 int

func main() {
	var num1 MyInt = 2
	var num2 MyInt = 6

	var num3 MyIntV2 = 3
	var num4 MyIntV2 = 4

	fmt.Println(sum(2, 4, 67))
	fmt.Println(sum[float32](2.0, 4.6, 67.1))
	fmt.Println(sum(num1, num2))
	fmt.Println(sum(num3, num4))
}

func sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}

	return total
}
