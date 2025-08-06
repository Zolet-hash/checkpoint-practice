package main

import (
	"fmt"
)

func ToLower(s string) string {
	result := ""

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r + ('a' - 'A') // Convert uppercase to lowercase
		}
		result += string(r)
	}

	return result
}

func main() {
	fmt.Println(ToLower("HELLO"))        // "hello"
	fmt.Println(ToLower("Hello WORLD!")) // "hello world!"
	fmt.Println(ToLower("Golang123"))    // "golang123"
	fmt.Println(ToLower(""))             // ""

}
