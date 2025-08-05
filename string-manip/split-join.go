//unfinished
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a,b,c,d"
	parts := strings.split(s, ",")
	fmt.Println((parts))

	joined := strings.join(parts, "_")
	fmt,Println(joined)
}
