package main

import (
	"fmt"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(IsUpper("HELLO"))     // true ✅
	fmt.Println(IsUpper("HELLO!"))    // false ❌ ('!' is not a letter)
	fmt.Println(IsUpper(""))          // true ✅ (empty string = no violation)
	fmt.Println(IsUpper("UPPERCASE")) // true ✅
	fmt.Println(IsUpper("Upper"))     // false ❌ ('p' is lowercase)
}
