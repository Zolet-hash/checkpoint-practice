package main

import (
	"fmt"
)

func IsSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func main() {
	fmt.Println(IsSpace(' '))
	fmt.Println(IsSpace('\t'))
	fmt.Println(IsSpace('a'))
}
