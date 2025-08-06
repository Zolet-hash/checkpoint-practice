package main

import (
	"fmt"
)

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello!"))    // true ✅
	fmt.Println(IsPrintable("123 @#"))    // true ✅
	fmt.Println(IsPrintable("Hi\nThere")) // false ❌ (newline)
	fmt.Println(IsPrintable("你好"))        // false ❌ (Unicode)
	fmt.Println(IsPrintable(""))          // true ✅ (empty = all valid)

}
