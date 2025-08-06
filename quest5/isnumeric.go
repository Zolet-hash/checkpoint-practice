package main

import (
	"fmt"
)

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(piscine.IsNumeric("12345")) // true ✅
	fmt.Println(piscine.IsNumeric("000"))   // true ✅
	fmt.Println(piscine.IsNumeric("12a45")) // false ❌
	fmt.Println(piscine.IsNumeric("12 45")) // false ❌
	fmt.Println(piscine.IsNumeric(""))      // true ✅

}
