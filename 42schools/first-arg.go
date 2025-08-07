package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there are any arguments passed (excluding the program name)
	if len(os.Args) < 2 {
		// If no arguments are passed, print an empty line
		fmt.Println()
	} else {
		// Otherwise, print the first argument
		fmt.Println(os.Args[1])
	}
}

/*Write a program that takes strings as arguments, and displays its first
argument followed by a \n.

If the number of arguments is less than 1, the program displays \n.

Example:

$> ./aff_first_param vincent mit "l'ane" dans un pre et "s'en" vint | cat -e
vincent$
$> ./aff_first_param "j'aime le fromage de chevre" | cat -e
j'aime le fromage de chevre$
$> ./aff_first_param
$*/