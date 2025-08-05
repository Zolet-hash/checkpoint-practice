/*package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.iterativefactorial(args))

}

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	return nb * IterativeFactorial(nb-1)
}
*/

package main

import (
	"fmt"
	"math"
)

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0 //handles -ve  value{-ve values lacks factorials}
	}
	result := 1

	for i := 2; i <= nb; i++ {
		//check for overflow before multiplication
		if result > math.MaxInt/int(i) {
			return 0 //overflow will happen here
		}
		result *= i
	}
	return result
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

/*iterativefactorial
Instructions
Write an iterative function that returns the factorial of the int passed as parameter.

Errors (non possible values or overflows) will return 0.
*/
