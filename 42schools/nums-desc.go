//Write a program that displays all digits in descending order,
// followed by a newline

package main

import (
	"fmt"
)

func main() {
	for n := 9; n >= 0; n-- {
		fmt.Print(n) //no new line at the end

	}

	fmt.Println() //add new line after all digits
}
