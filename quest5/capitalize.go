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
				runes[i] = r - ('a' - 'A') // lowercase to uppercase
			} else if !capitalizeNext && isUpper(r) {
				runes[i] = r + ('a' - 'A') // uppercase to lowercase
			}
			capitalizeNext = false
		} else {
			capitalizeNext = true
		}
	}
	return string(runes)
}

// Helpers
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

func main() {
	fmt.Println(Capitalize("hello world!"))        // "Hello World!"
	fmt.Println(Capitalize("42school is AWESOME")) // "42school Is Awesome"
	fmt.Println(Capitalize("goLang123language"))   // "Golang123language"
	fmt.Println(Capitalize("HELLO, my FRIEND!"))   // "Hello, My Friend!"

}

/*💡 How It Works

  Converts the input string into a rune slice to allow modification.

  Uses a flag capitalizeNext to decide if the next alphanumeric rune should be capitalized.

  If it's the start of a word and the letter is lowercase → uppercase it.

  If it's the middle of a word and the letter is uppercase → lowercase it.

  Resets capitalizeNext = true when encountering a non-alphanumeric character.*/
