package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {

	errorFormatter := func(first, second string) string {
		return fmt.Sprintf("Result should be %v but got %v instead", first, second)
	}

	t.Run("Reverse reverses a string", func(t *testing.T) {
		word := "abcd"
		result := reverseStr(word)
		expect := "dcba"

		if result != expect {
			t.Error(errorFormatter(expect, result))
		}
	})

	t.Run("Reverse reverses a string", func(t *testing.T) {
		word := "   abcd"
		result := reverseStr(word)
		expect := "dcba   "

		if result != expect {
			t.Error(errorFormatter(expect, result))
		}
	})
}
