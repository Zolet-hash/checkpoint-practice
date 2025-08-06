package main

import (
	"fmt"
)

// Fibonacci returns the nth Fibonacci number using recursion.
// Index starts at 0: Fibonacci(0) = 0, Fibonacci(1) = 1, Fibonacci(2) = 1, ...
// Negative indices return -1.
func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}

func main() {
	arg1 := 4

	fmt.Println(Fibonacci(arg1))
	fmt.Println()
}
