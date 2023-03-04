package main

import "fmt"

func main() {
	division(100, 10)
	division(200, 25)
	division(34, 0)
	division(124, 8)
}

func division(dividend, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("me recupere del panic")
		}
	}()

	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("🤕 no puedes dividir por cero")
	}
}
