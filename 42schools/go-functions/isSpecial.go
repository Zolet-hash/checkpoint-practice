package main

import (
	"fmt"
	"unicode"
)

func IsSpecial(s string) int {
	count := 0

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(IsSpecial("He@#llo ^&td"))
	fmt.Println(IsSpecial("H78"))
}
