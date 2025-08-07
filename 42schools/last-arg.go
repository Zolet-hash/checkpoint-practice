package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there are no arguments (program name doesn't count)
	if len(os.Args) < 2 {
		fmt.Println() // Print just a newline if no arguments are given
		return
	}

	// Get the last argument (os.Args[0] is the program name)
	lastArg := os.Args[len(os.Args)-1]
	
	// Print the last argument followed by a newline
	fmt.Println(lastArg)
}

/*Write a program that takes strings as arguments, and displays its last
argument followed by a newline.

If the number of arguments is less than 1, the program displays a newline.*/