package main

import (
	"fmt"
)

func Compare(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)

	minLen := len(runesA)
	if len(runesB) < minLen {
		minLen = len(runesB)
	}

	for i := 0; i < minLen; i++ {
		if runesA[i] < runesB[i] {
			return -1
		} else if runesA[i] > runesB[i] {
			return 1
		}
	}

	// If all compared runes are equal, the shorter string is "less"
	if len(runesA) < len(runesB) {
		return -1
	} else if len(runesA) > len(runesB) {
		return 1
	}

	return 0
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
