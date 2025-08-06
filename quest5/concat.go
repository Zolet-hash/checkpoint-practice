package main

import (
	"fmt"
)

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

package main() {
	fmt.Println(Concat("Hello!", " How are you?"))
	// Output: Hello! How are you?
	
}
