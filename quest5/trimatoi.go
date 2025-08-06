package main

import (
	"fmt"
)

func main() {
	fmt.Println(TrimAtoi("12345"))        // 12345
	fmt.Println(TrimAtoi("str123ing45"))  // 12345
	fmt.Println(TrimAtoi("012 345"))      // 12345
	fmt.Println(TrimAtoi("Hello World!")) // 0
	fmt.Println(TrimAtoi("sd+x1fa2W3s4")) // 1234
	fmt.Println(TrimAtoi("sd-x1fa2W3s4")) // -1234
	fmt.Println(TrimAtoi("sdx1-fa2W3s4")) // 1234
	fmt.Println(TrimAtoi("sdx1+fa2W3s4")) // 1234
}

func TrimAtoi(s string) int {
	num := 0
	negative := false
	foundDigit := false

	for _, r := range s {
		if r == '-' && !foundDigit {
			negative = true
		}
		if r >= '0' && r <= '9' {
			foundDigit = true
			num = num*10 + int(r-'0')
		}
	}

	if !foundDigit {
		return 0
	}

	if negative {
		return -num
	}

	return num
}
