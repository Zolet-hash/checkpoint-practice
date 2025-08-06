package main

import "fmt"

func LastRune(s string) rune {
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}

func main() {
	fmt.Printf("%c\n", LastRune("Emmanuel")) // Output: l
	fmt.Printf("%c\n", LastRune("GoLang🚀"))  // Output: 🚀
	fmt.Printf("%c\n", LastRune("¡Hola!"))   // Output: !
}

/*ast = r stores each rune as we loop.

When the loop ends, last holds the last rune in the string.

Handles multibyte characters like emojis and accented letters correctl*/

