// --- Directions
// Check to see if two provided strings are anagrams of eachother.
// One string is an anagram of another if it uses the same characters
// in the same quantity. Only consider characters, not spaces
// or punctuation.  Consider capital letters to be the same as lower case
// --- Examples
//
//	anagrams('rail safety', 'fairy tales') --> True
//	anagrams('RAIL! SAFETY!', 'fairy tales') --> True
//	anagrams('Hi there', 'Bye there') --> False
package main

import (
	"fmt"
	"unicode"
)

func main() {
	firstRes := anagrams("A tree, a life, a bench", "A tree, a fence, a yard")
	fmt.Printf("First result is %t \n", firstRes)
}

func anagrams(firstWord, secondWord string) bool {
	charArr := make([]int, 26, 26)
	specialChar := map[byte]bool{
		' ': true,
		',': true,
		'!': true,
	}

	if len(firstWord) != len(secondWord) {
		return false
	}

	for i := 0; i < len(firstWord); i++ {
		if _, ok := specialChar[firstWord[i]]; !ok {
			firstWordChar := unicode.ToLower(rune(firstWord[i])) - 'a'
			charArr[firstWordChar]++
		}

		if _, ok := specialChar[secondWord[i]]; !ok {
			secondWordChar := unicode.ToLower(rune(secondWord[i])) - 'a'
			charArr[secondWordChar]--
		}
	}

	for _, value := range charArr {
		if value > 0 {
			return false
		}
	}
	return true
}
