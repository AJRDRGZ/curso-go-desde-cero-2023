package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Includes([]string{"a", "b", "c"}, "c"))
	fmt.Println(Includes([]string{"a", "b", "c"}, "d"))
	fmt.Println(Includes([]int{1, 12, 24}, 24))

	fmt.Println(filter([]int{2, 12, 23, 98, 21, 79}, lessThanTwenty))
}

func lessThanTwenty(num int) bool {
	return num < 20
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func filter[T constraints.Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}
