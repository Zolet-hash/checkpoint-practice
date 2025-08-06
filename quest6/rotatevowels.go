package main

import (
	"fmt"
	"os"
	"strings"
)

func isVowel(r rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, r)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println()
		return
	}

	// Join all args with spaces
	input := strings.Join(args, " ")

	// Collect vowels
	var vowels []rune
	for _, ch := range input {
		if isVowel(ch) {
			vowels = append(vowels, ch)
		}
	}

	// If no vowels, print the input and return
	if len(vowels) == 0 {
		fmt.Println(input)
		return
	}

	// Reverse vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Replace vowels in input
	result := []rune(input)
	vowelIndex := 0
	for i, ch := range result {
		if isVowel(ch) {
			result[i] = vowels[vowelIndex]
			vowelIndex++
		}
	}

	fmt.Println(string(result))
}
