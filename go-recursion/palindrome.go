//recursive palindrome
//word that reads the same backword & forward ie madam racecar

package main

import("fmt")


//function to check if a word is a palindrome 
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true // base case, empty or one-letter case are palindromes
	}
	if s[0] != s[len(s)-1] {
		return false //mostmatch at the ends: not a palindrome
	}
	return isPalindrome(s[1 : len(s)-1]) //recurse with inner substring
}

func main() {
	word := "madam"
	if isPalindrome(word) {
		fmt.Println(word, "is a palindrome")
	} else {
		fmt.Println(word, "is not a palindrome")
	}
}