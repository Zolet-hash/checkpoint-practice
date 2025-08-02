//Write a function that takes a pointer to a pointer to a pointer to an int as argument and gives to this int the value of 1.

package main

import (
	"fmt"
)

func main() {  //ignore the vs code eror
	a := 0
	b := &a
	n := &b
	ultimatePointer(&n)

	fmt.Println(a)

}

func ultimatePointer(a ***int) {
	***a = 1
}
