package main

import (
	"fmt"
)

func CountAlpha(s string) int {

	count := 0

	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			count++
		}

	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello World"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1l2l3o0"))
}
