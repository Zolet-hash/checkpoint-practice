package main

// Write a function that takes a pointer to an int as argument and gives to this int the value of 1.

//Expected function
//func PointOne(n *int) {

//}
import (
	"fmt"
)

func main() {
	n := 0

	pointOne(&n) //pass its address to pointone
	fmt.Println(n)

}

func pointOne(n *int) {
	*n = 1
}
