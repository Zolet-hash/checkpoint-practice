package main

import (
	"fmt"
)

func ASCIIToInt(n byte) int {
	return (int(n))
}

func main() {
	fmt.Println(ASCIIToInt('A'))
}
