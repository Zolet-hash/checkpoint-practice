package main

import "fmt"

func FirstRune(s string) rune {
	for _, r := range s {
		return r // return the first rune we encounter
	}
	return 0 // if string is empty, return 0
}

func main() {
	fmt.Printf("%c\n", FirstRune("Emmanuel")) // Output: E
	fmt.Printf("%c\n", FirstRune("🦁Lion"))    // Output: 🦁
	fmt.Printf("%c\n", FirstRune("¡Hola!"))   // Output: ¡
}
