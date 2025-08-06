package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:] // Skip the program name

	sort.Strings(args) // Sort in ASCII order

	for _, arg := range args {
		fmt.Println(arg)
	}
}

//INCOMPLETE!!

/*Instructions
Write a program that prints the arguments received in the command line in ASCII order.

Example of output :

$ go run . 1 a 2 A 3 b 4 C
1
2
3
4
A
C
a
b
$*/
