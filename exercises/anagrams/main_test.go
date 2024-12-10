package main

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	errorFormat := func(firstArg, secondArg bool) string {
		return fmt.Sprintf("Result should be %v but got %v instead", firstArg, secondArg)
	}

	// Import anagram function
	t.Run(" \"Hello\" is an anagrams of \"llohe\"", func(t *testing.T) {
		firstWord := "Hello"
		secondWord := "llohe"
		result := anagrams(firstWord, secondWord)
		expect := true

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run(" \"Whoa! Hi!\" is an anagram of \"Hi! Whoa!\"", func(t *testing.T) {
		firstWord := "Whoa! Hi!"
		secondWord := "Hi! Whoa!"
		result := anagrams(firstWord, secondWord)
		expect := true

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"One One\" is not an anagram of \"Two two two\"", func(t *testing.T) {
		firstWord := "One One"
		secondWord := "Two two two"

		result := anagrams(firstWord, secondWord)
		expect := false

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"One one\" is not an anagram of \"One one c\"", func(t *testing.T) {
		firstWord := "One one"
		secondWord := "One one c"

		result := anagrams(firstWord, secondWord)
		expect := false

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"A tree, a life, a bench\" is not an anagram of \"A tree, a fence, a yard\"", func(t *testing.T) {
		firstWord := "A tree, a life, a bench"
		secondWord := "A tree, a fence, a yard"

		result := anagrams(firstWord, secondWord)
		expect := false

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})
}
