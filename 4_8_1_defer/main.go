package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello EDteam"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
