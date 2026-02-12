// factorial of a number
package main
import ("fmt")

func factorial(n int) int {
	//base case
		if n == 0 {
			return 1
		}
	//calling factoral func recurively until it reaches the base case
		return n * factorial(n-1)
}

func main() {
		fmt.Println(factorial(4))
}
