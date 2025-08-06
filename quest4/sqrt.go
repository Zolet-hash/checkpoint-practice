package main

import "fmt"

// Sqrt returns the square root of nb if it's a perfect square.
// If nb is not a perfect square, it returns 0.
func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	i := 0
	for i*i <= nb {
		if i*i == nb {
			return i
		}
		i++
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(25)) // 5
	fmt.Println(Sqrt(10)) // 0
	fmt.Println(Sqrt(4))  // 1
	fmt.Println(Sqrt(3))  // 0
	fmt.Println(Sqrt(-9)) // 0
}
