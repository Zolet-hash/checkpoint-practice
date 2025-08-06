package main

import (
	"fmt"
)

func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))  // true ✅
	fmt.Println(IsLower("hello!")) // false ❌
	fmt.Println(IsLower("Hello"))  // false ❌
	fmt.Println(IsLower("world"))  // true ✅
	fmt.Println(IsLower(""))       // true ✅ (empty string = no violation)

}
