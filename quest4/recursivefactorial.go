package main

import (
	"fmt"
)

func RecursiveFactorial(nb int) int {
	if nb <= 0 { //handling the -ve efficiently
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}

/*Instructions
Write a recursive function that returns the factorial of the int passed as parameter.

Errors (non possible values or overflows) will return 0.

for is forbidden for this exercise. */
