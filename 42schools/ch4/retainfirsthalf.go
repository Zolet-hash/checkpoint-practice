package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
    length := len(str) 
    
    if length == 0 || length == 1 {
        return str
    }
    
    half := length/2
    return str[:half]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

