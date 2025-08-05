//reverse a string using recursion

package main

import ("fmt")

func main() {
	str := "sulwe"

	reversed := reverse(str)
	fmt.Println("Reversed:", reversed)
}

func reverse(s string) string {  //a function named reverse that takes a string 's' & returns a string
	if len(s) == 0 {     //base of the recursion,, if the string is empty, we return an empty string and stop the recursion 
		return ""
	}
	return reverse(s[1:]) + string(s[0]) 
}