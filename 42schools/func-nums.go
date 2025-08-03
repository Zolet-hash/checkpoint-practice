// Write a function that displays all digits in descending order.
package main

import (
	"fmt"
)

func funcNums() {
	for n := 9; n >= 0; n-- {
		fmt.Print(n)
	}
}
func main() {

	funcNums()    //parsing our func to main func
	fmt.Println() //printing an empty line
}
