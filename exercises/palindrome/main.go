package main

import "fmt"

func main() {
	test := "pennep"
	fmt.Println(palindrome(test))
}

func palindrome(str string) bool {
	firstPointer := 0
	secondPointer := len(str) - 1

	for firstPointer < secondPointer {
		if str[firstPointer] != str[secondPointer] {
			return false
		}
		firstPointer++
		secondPointer--
	}
	return true
}
