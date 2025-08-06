package main

import "fmt"

// IsPrime returns true if nb is a prime number, false otherwise.
func IsPrime(nb int) bool {
	if nb <= 1 { //prime numbers must be greater than 1
		return false
	}
	if nb == 2 { //it's the only prime number
		return true
	}
	if nb%2 == 0 { //any other even number is nor a prime
		return false
	}

	// Check from 3 up to square root of nb, skipping even numbers
	for i := 3; i*i <= nb; i += 2 { //Loop from 3 up to the square root of nb, skipping even numbers (i += 2) If any number divides nb, it's not prime.
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(2))   // true
	fmt.Println(IsPrime(3))   // true
	fmt.Println(IsPrime(4))   // false
	fmt.Println(IsPrime(17))  // true
	fmt.Println(IsPrime(1))   // false
	fmt.Println(IsPrime(997)) // true
}
