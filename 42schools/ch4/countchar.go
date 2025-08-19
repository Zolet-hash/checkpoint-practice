package main

import (
	"fmt"
)

func CountChar(str string, c rune) int {
	if str == "" {
		return 0
	}

	count := 0
	for _, r := range str {
		if r == c {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
