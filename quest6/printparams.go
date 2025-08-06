package main

import (
	"fmt"
	"os"
)

func main() {
	// Start from index 1 to skip the program name (os.Args[0])
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}

//to see it -> go run . choumi is the best cat
