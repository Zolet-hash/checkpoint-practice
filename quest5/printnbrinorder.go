package main

import "fmt"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	// Special case for 0
	if n == 0 {
		fmt.Print(0)
		return
	}

	// Extract digits
	var digits []int
	for n > 0 {
		d := n % 10
		digits = append(digits, d)
		n = n / 10
	}

	// Sort digits (simple bubble sort since we can't import sort package)
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[j] < digits[i] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	// Print sorted digits
	for _, d := range digits {
		fmt.Print(d)
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
