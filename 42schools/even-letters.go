//Write a program that displays the alphabet, with even letters in uppercase, and
//odd letters in lowercase, followed by a newline.

package main

import (
	"fmt"
	"unicode"
)

func main() {
	index := 0

	for ch := 'a'; ch <= 'z'; ch++ {
		if index%2 == 0 {
			fmt.Printf("%c", ch)
		} else {
			fmt.Printf("%c", unicode.ToUpper(ch))
		}
		index++
	}
	fmt.Println()
}
