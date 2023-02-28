package main

import "fmt"

func main() {
	music := make(map[string]string)
	music["guitar"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	tech := map[string]string{
		"computer": "💻",
		"mouse":    "🖱️",
	}

	fmt.Println(tech)

	delete(tech, "computer")
	fmt.Println(tech)

	content, ok := music["guitar"]

	fmt.Println(content, ok)
}
