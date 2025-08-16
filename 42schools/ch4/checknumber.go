package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
    
    numbers := "0123456789"
    
    for _, r := range arg {
        for _, n := range numbers {
            if r == n {
                return true
            }
        }
    }
    return false
    
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}


//Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.
