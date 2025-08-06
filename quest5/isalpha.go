package main

import (
	"fmt"
)

func IsAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello123"))  // true ✅
	fmt.Println(IsAlpha("Hello 123")) // false ❌ (space is invalid)
	fmt.Println(IsAlpha("Hi!"))       // false ❌ (! is invalid)
	fmt.Println(IsAlpha("42"))        // true ✅
	fmt.Println(IsAlpha(""))          // true ✅ (empty = valid)

}
