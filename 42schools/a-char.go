package main

import (
	"fmt"
	"os"
)

func main() {
	//no of args !=2(program name + 1 argument)
	if len(os.Args) != 2 {
		fmt.Println("a")
		return //exits the program immedeately
	}

	//extract the single arguments and store it inside a variable
	input := os.Args[1]

	for _, char := range input {
		if char == 'a' {
			fmt.Println("a")
			return //call os.Status exit[1]
		}
	}

	fmt.Println()
}
