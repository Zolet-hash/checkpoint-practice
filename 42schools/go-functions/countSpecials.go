package main

import (
	"fmt"
)

func CountSpecials(s string) int {

	specials := "!@#$%^&*()_+="
	count := 0

	for _, r := range s {
		for _, sp := range specials {
			if r == sp {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(CountSpecials("Hel#$#$llo W&&ld"))
}
