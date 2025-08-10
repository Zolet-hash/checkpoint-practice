package main

import (
	"fmt"
)

func ToUpper(s string) string {
	result := []rune(s) //converts to rune for easy manipulation

	for i, r := range result {
		if r >= 'a' && r <= 'z' {
			result[i] = r - 32
		}
		if r >= 'A' && r <= 'Z' {
			result[i] = r + 32
		}
	}
	return string(result)
}

func main() {
	fmt.Println(ToUpper("heLlo world"))
	fmt.Println(ToUpper("GOLANg is powerful"))
}
