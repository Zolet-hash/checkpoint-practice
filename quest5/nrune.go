package main

import "fmt"

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	index := 1
	for _, r := range s {
		if index == n {
			return r
		}
		index++
	}
	return 0 // n is out of bounds
}

func main() {
	fmt.Printf("%c\n", NRune("Emmanuel", 1))  // Output: E
	fmt.Printf("%c\n", NRune("Emmanuel", 3))  // Output: m
	fmt.Printf("%c\n", NRune("¡Hola!", 2))    // Output: H
	fmt.Printf("%c\n", NRune("🚀Rocket", 1))   // Output: 🚀
	fmt.Printf("%c\n", NRune("🚀Rocket", 100)) // Output: (empty)
}
