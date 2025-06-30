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
	runes := []rune(cleanText)
	length := len(cleanText)
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func reversedString(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countVowels(text string) int {
	var count int = 0
	lowerText := strings.ToLower(text)
	runes := []rune(lowerText)
	for i := 0; i < len(runes); i++ {
		switch runes[i] {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
