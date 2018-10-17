package main

import (
	"fmt"
)

func main() {
	colors := map[int]string{
		1: "red",
		2: "green",
		3: "purple",
	}
	fmt.Println(colors)

	for key, value := range colors {
		fmt.Println(key, value)
	}
}
