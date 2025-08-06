package main

import (
	"fmt"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems)) // Output: Hello! How are you?
}

func BasicJoin(elems []string) string {
	result := ""
	for _, str := range elems {
		result += str
	}
	return result
}
