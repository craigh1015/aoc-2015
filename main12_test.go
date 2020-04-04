package main

import "testing"

func TestGetNumbersTotalJSON12(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected float64
	}{
		{desc: "01", input: `[]`, expected: 0},
		{desc: "02", input: `{}`, expected: 0},
		{desc: "03", input: `[1,2,3]`, expected: 6},
		{desc: "04", input: `{"a":2,"b":4}`, expected: 6},
		{desc: "05", input: `[[[3]]]`, expected: 3},
		{desc: "06", input: `{"a":{"b":4},"c":-1}`, expected: 3},
		{desc: "07", input: `{"a":[-1,1]}`, expected: 0},
		{desc: "08", input: `[-1,{"a":1}]`, expected: 0},
		{desc: "09", input: `[1,{"c":"red","b":2},3]`, expected: 4},
		{desc: "10", input: `{"d":"red","e":[1,2,3,4],"f":5}`, expected: 0},
		{desc: "11", input: `[1,"red",5]`, expected: 6},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := getNumbersTotalJSON12(tC.input)
			if actual != tC.expected {
				t.Fatalf("%s - Expected: %v got: %v\n", tC.input, tC.expected, actual)
			}
		})
	}
}
