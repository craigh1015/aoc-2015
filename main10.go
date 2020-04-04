package main

import (
	"bytes"
	"strconv"
)

func main10() {
}

func lookAndSay10(input string, iterations int) string {
	result := input
	for i := 0; i < iterations; i++ {
		result = iterate(result)
	}
	return result
}

func iterate(input string) string {
	var buffer bytes.Buffer
	count := 0
	current := 'X'
	for _, char := range input {
		if count == 0 {
			current = char
		}
		if char == current {
			count++
			continue
		}
		buffer.WriteString(strconv.Itoa(count))
		buffer.WriteString(string(current))
		current = char
		count = 1
	}
	buffer.WriteString(strconv.Itoa(count))
	buffer.WriteString(string(current))
	return buffer.String()
}
