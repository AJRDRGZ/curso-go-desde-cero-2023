package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("not found")

var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

func main() {
	found, err := search("34")
	if errors.Is(err, errNotFound) {
		fmt.Println("pudimos controlar el error correctamente")
		return
	}
	if err != nil {
		fmt.Println("search()", err)
		return
	}

	fmt.Println(found)
}

func search(key string) (string, error) {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %w", err)
	}

	emoji, err := findFood(num)
	if err != nil {
		return "", fmt.Errorf("findFood(): %w", err)
	}

	return emoji, nil
}

func findFood(id int) (string, error) {
	value, ok := food[id]
	if !ok {
		return "", errNotFound
	}

	return value, nil
}
