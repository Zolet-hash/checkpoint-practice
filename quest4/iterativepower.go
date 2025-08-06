// Iteratative power returns nb to the power of power
package main

import (
	"fmt"
)

func IterativePower(nb int, power int) int {

	if power < 0 {
		return 0 // -ve exponents not handled
	}
	result := 1
	for i := 0; i <= power; i++ {
		result *= nb
		//result *= nb //shorthand for result = result *nb
	}
	return result
}

func main() {
	fmt.Println(IterativePower(4, 3))
	fmt.Println()

}
