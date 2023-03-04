package main

import "fmt"

func main() {
	nums := []int{2, 12, 23, 98, 21, 79}

	result := filter(nums, lessThanTwenty)
	fmt.Println(result)

	fmt.Println(sum(2)(3))

	plus10 := sum(10)
	fmt.Println(plus10(2))
	fmt.Println(plus10(4))
	fmt.Println(plus10(1))
	fmt.Println(plus10(5))
}

func sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func greatherToTen(num int) bool {
	return num > 10
}

func lessThanTwenty(num int) bool {
	return num < 20
}

func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}
