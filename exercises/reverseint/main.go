package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverseInt(-900))
}

func reverseInt(num int) int {
	var isMinusInt bool
	var buffer bytes.Buffer
	numStr := strconv.FormatInt(int64(num), 10)
	if num < 0 {
		isMinusInt = true
		numStr = numStr[1:]
	}
	for i := len(numStr) - 1; i >= 0; i-- {
		buffer.WriteString(string(numStr[i]))
	}
	var zeroCnt int
	resultFromBuffer := buffer.String()
	for _, char := range resultFromBuffer {
		if char == '0' {
			zeroCnt++
		} else {
			break
		}
	}
	var resultStr string
	if isMinusInt {
		resultStr = "-"
	}
	resultStr += resultFromBuffer[zeroCnt:]
	result, _ := strconv.Atoi(resultStr)
	return result
}
