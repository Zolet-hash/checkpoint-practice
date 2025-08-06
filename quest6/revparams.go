package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Skip the program name

	// Loop in reverse
	for i := len(args) - 1; i >= 0; i-- {
		fmt.Println(args[i])
	}
}

/* Write a program that prints the arguments received in the command line in reverse order.

Example of output :

$ go run . choumi is the best cat
cat
best
the
is
choumi
*/
