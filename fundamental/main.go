package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Hannah"
	age := 10
	printFullName(age, name)
	fmt.Println()
	if isPalindrome(name) {
		fmt.Printf("Is palindrome\n")
	} else {
		fmt.Printf("Not palindrome\n")
	}

	fmt.Printf("The reversed string of %s is %s\n", name, reversedString(name))
	fmt.Printf("The vowels count is %d", countVowels(name))
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

func reversedString(text string) string {
	var reversedString string
	for i := len(text); i > 0; i-- {
		reversedString += string(text[i-1])
	}
	return reversedString
}

func countVowels(text string) int {
	var count int = 0
	lowerText := strings.ToLower(text)
	for i := 0; i < len(lowerText); i++ {
		switch lowerText[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
