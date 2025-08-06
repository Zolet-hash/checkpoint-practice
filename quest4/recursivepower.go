package main

import (
	"fmt"
)

func RecursivePower(nb int, pow int) int {
	if pow < 0 {
		return 0
	}
	if pow == 0 {
		return 1 //base case: any number to power 0 is 1
	}
	return nb * RecursivePower(nb, pow-1) //this is the recursive case, we reduce the problem -1
}

func main() {
	fmt.Println(RecursivePower(4, 3))
}

/* his is the recursive case:

   We reduce the problem by one step each time.

   Example: RecursivePower(4, 3) → 4 * RecursivePower(4, 2)

   Then: 4 * (4 * RecursivePower(4, 1))

   Then: 4 * (4 * (4 * RecursivePower(4, 0)))

   Finally: 4 * 4 * 4 * 1 = 64*/
