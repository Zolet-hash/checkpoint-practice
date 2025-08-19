package main

import (
	"fmt"
	
)

func FishAndChips( n int) string {
    if n < 0 {
        return "error: number is negative"
    }
    
    divBy2 := n%2 == 0
    divBy3 := n%3 == 0
    
    switch {
        case divBy2 && divBy3:
        return "fish and chips"
        
        case divBy2:
        return "fish"
        
        case divBy3:
        return "chips"
        
        default:
        return "error: not divisible"
    }
}

func main() {
	fmt.Println(FishAndChips(4))  // divisible by 2
	fmt.Println(FishAndChips(9))  // divisible by 3
	fmt.Println(FishAndChips(6))  // divisible by 2 and 3
}
