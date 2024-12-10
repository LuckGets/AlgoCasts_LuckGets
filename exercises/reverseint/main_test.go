package main

import (
	"fmt"
	"testing"
)

func TestReverseInt(t *testing.T) {
	errorFormatter := func(firstArg, secondArg int) string {
		return fmt.Sprintf("Result should be %d but got %d instead", firstArg, secondArg)
	}

	t.Run("it should handles 0 as an input", func(t *testing.T) {
		num := 0
		res := reverseInt(num)
		expect := 0

		if res != expect {
			t.Error(errorFormatter(expect, res))
		}
	})
}
