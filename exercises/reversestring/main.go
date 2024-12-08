package main

import (
	"bytes"
	"fmt"
)

func main() {
	test := "apple"
	fmt.Println(reverseStr(test))
}

func reverseStr(word string) string {
	var res bytes.Buffer
	for i := len(word) - 1; i >= 0; i-- {
		res.WriteString(string(word[i]))
	}
	return res.String()
}
