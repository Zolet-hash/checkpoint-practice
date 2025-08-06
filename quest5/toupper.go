package main

import (
	"fmt"
)

func ToUpper(s string) string {
	result := ""

	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = r - ('a' - 'A') // Convert lowercase to uppercase
		}
		result += string(r)
	}

	return result
}

func main() {
	fmt.Println(ToUpper("hello"))      // "HELLO"
	fmt.Println(ToUpper("Hello 123!")) // "HELLO 123!"
	fmt.Println(ToUpper("golang!"))    // "GOLANG!"
	fmt.Println(ToUpper(""))           // ""

}
