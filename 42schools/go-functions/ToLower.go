package main

import (
	"fmt"
)

func ToLower(s string) string {
	result := []rune(s)

	for i, r := range result {
		if r >= 'A' && r <= 'z' {
			result[i] = r + 32
		}
	}
	return string(result)
}

func main() {
	fmt.Println(ToLower("HELLO WORLD"))
}
