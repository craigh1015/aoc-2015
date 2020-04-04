package main

import (
	"log"
	"testing"
)

func TestLookAndSay10(t *testing.T) {
	t.Skip()
	testCases := []struct {
		desc       string
		input      string
		iterations int
		expected   string
	}{
		{desc: "01", input: "1", iterations: 1, expected: "11"},
		{desc: "02", input: "11", iterations: 1, expected: "21"},
		{desc: "03", input: "21", iterations: 1, expected: "1211"},
		{desc: "04", input: "1211", iterations: 1, expected: "111221"},
		{desc: "05", input: "111221", iterations: 1, expected: "312211"},
		{desc: "06", input: "1", iterations: 5, expected: "312211"},
		{desc: "07", input: "3113322113", iterations: 40, expected: ""},
		{desc: "08", input: "3113322113", iterations: 50, expected: ""},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := lookAndSay10(tC.input, tC.iterations)
			log.Printf("input: %s -> output len: %d\n", tC.input, len(actual))
			if len(tC.expected) > 0 && actual != tC.expected {
				t.Fatalf("Expected: %s got: %s\n", tC.expected, actual)
			}
		})
	}
}
