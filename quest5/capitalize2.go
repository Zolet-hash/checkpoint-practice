package main

import (
	"fmt"
)

func Capitalize(s string) string {
	runes := []rune(s)
	capitalizeNext := true

	for i, r := range runes {
		if isAlphaNumeric(r) {
			if capitalizeNext && isLower(r) {
				runes[i] = r - ('a' - 'A')
			} else if !capitalizeNext && isUpper(r) {
				runes[i] = r + ('a' - 'A')
			}
			capitalizeNext = false
		} else {
			capitalizeNext = true
		}
	}

	return string(runes)
}

// Helper functions

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
