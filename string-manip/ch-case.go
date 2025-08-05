//change case

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "emmanuel sulwe"

	fmt.Println("CapsLock:", strings.ToUpper(s))
	fmt.Println("lowercase:", strings.ToLower(s))

	fmt.Println(strings.Title(s))
}
