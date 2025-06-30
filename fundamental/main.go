package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Go lang"
	age := 10
	printFullName(age, name)
	fmt.Println()
	if isPalindrome("bab") {
		fmt.Printf("Is palindrome")
	} else {
		fmt.Printf("Not palindrome")
	}
}

func printFullName(age int, name string) {
	fmt.Printf("Name is %s and %d years old.", name, age)
}

func isPalindrome(text string) bool {
	cleanText := strings.ToLower(strings.ReplaceAll(text, " ", ""))
	length := len(cleanText)
	for i := 0; i < length/2; i++ {
		if cleanText[i] != cleanText[length-1-i] {
			return false
		}
	}
	return true
}
