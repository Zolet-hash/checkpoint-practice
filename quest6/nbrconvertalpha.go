package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	upper := false
	startIndex := 0

	// Check for the --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		startIndex = 1
	}

	// Process arguments starting from the correct index
	for _, arg := range args[startIndex:] {
		n, err := strconv.Atoi(arg)
		if err != nil || n < 1 || n > 26 {
			fmt.Print(" ")
			continue
		}

		letter := 'a' + rune(n-1)
		if upper {
			letter = 'A' + rune(n-1)
		}
		fmt.Print(string(letter))
	}
	fmt.Println()
}

/* Example Usage
Lowercase:

$ go run . 8 5 12 12 15
hello

Uppercase:

$ go run . --upper 8 5 25
HEY

Invalid args:

$ go run . 32 86 h
   # (prints 3 spaces)*/
