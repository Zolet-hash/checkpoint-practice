package main

import "fmt"

// IsPrime checks if a number is prime (optimized)
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// FindNextPrime returns the first prime number ≥ nb
func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	// If nb is even and > 2, make it odd
	if nb%2 == 0 {
		nb++
	}

	for {
		if IsPrime(nb) {
			return nb
		}
		nb += 2 // skip even numbers
	}
}

func main() {
	fmt.Println(FindNextPrime(14))  // 17
	fmt.Println(FindNextPrime(2))   // 2
	fmt.Println(FindNextPrime(0))   // 2
	fmt.Println(FindNextPrime(100)) // 101
}

/*IsPrime(nb int) bool

    Optimized to check divisibility only up to √nb.

    Skips even numbers after checking 2.

FindNextPrime(nb int) int

    Handles edge case: if nb ≤ 2, return 2 (first prime).

    If nb is even and > 2, increment it to start with an odd number (no even number >2 can be prime).

    Then loop, skipping even numbers, checking each until a prime is found.

*/
