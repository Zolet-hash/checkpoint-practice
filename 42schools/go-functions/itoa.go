package main

import (
	"fmt"
)

func IntToASCII(n int) string {
	return (string(rune(n)))
}

func main() {
	fmt.Println(IntToASCII(555))
	fmt.Println(IntToASCII(97))
	fmt.Println(IntToASCII(36))
}
