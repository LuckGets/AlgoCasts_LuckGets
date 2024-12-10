package main

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	errorFormat := func(firstArg, secondArg bool) string {
		return fmt.Sprintf("Result should be %v but got %v instead", firstArg, secondArg)
	}

	t.Run("\"aba\" is a palindrome", func(t *testing.T) {
		word := "aba"
		result := palindrome(word)
		expect := true

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"  aba\" is not a palindrome", func(t *testing.T) {
		word := "  aba"
		result := palindrome(word)
		expect := false

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"aba  \" is not a palindrome", func(t *testing.T) {
		word := "aba  "
		result := palindrome(word)
		expect := false

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"greetings  \" is not a palindrome", func(t *testing.T) {
		word := "greetings"
		result := palindrome(word)
		expect := false

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"1000000001  \" is a palindrome", func(t *testing.T) {
		word := "1000000001"
		result := palindrome(word)
		expect := true

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"Fish hsif\" is not a palindrome", func(t *testing.T) {
		word := "Fish hsif"
		result := palindrome(word)
		expect := false

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})

	t.Run("\"pennep\" is a palindrome", func(t *testing.T) {
		word := "pennep"
		result := palindrome(word)
		expect := true

		if result != expect {
			t.Error(errorFormat(expect, result))
		}
	})
}
