package main

import (
	"fmt"
)

func Index(s string, toFind string) int {
	sLen := len(s)
	findLen := len(toFind)

	if findLen == 0 || findLen > sLen {
		return -1
	}

	for i := 0; i <= sLen-findLen; i++ {
		if s[i:i+findLen] == toFind {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
