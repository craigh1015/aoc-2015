package main

import (
	"testing"
)

func TestValidateStraight11(t *testing.T) {
	t.Skip()
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "01", input: "aaaaaaaa", expected: false},
		{desc: "02", input: "hijklmmn", expected: true},
		{desc: "03", input: "aaaaaabc", expected: true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := validateStraight11(tC.input)
			if actual != tC.expected {
				t.Fatalf("%s - Expected: %v got: %v\n", tC.input, tC.expected, actual)
			}
		})
	}
}

func TestValidateLetters11(t *testing.T) {
	t.Skip()
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "01", input: "aaaiaaaa", expected: false},
		{desc: "02", input: "aaaoaaaa", expected: false},
		{desc: "03", input: "aaalaaaa", expected: false},
		{desc: "04", input: "aaaaaaaa", expected: true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := validateLetters11(tC.input)
			if actual != tC.expected {
				t.Fatalf("%s - Expected: %v got: %v\n", tC.input, tC.expected, actual)
			}
		})
	}
}

func TestValidatePairs11(t *testing.T) {
	t.Skip()
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "01", input: "aaaaaaaa", expected: false},
		{desc: "02", input: "aabcdeaa", expected: false},
		{desc: "03", input: "aabbaaaa", expected: true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := validatePairs11(tC.input)
			if actual != tC.expected {
				t.Fatalf("%s - Expected: %v got: %v\n", tC.input, tC.expected, actual)
			}
		})
	}
}

func TestValidate11(t *testing.T) {
	t.Skip()
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{desc: "01", input: "hijklmmn", expected: false},
		{desc: "02", input: "abbceffg", expected: false},
		{desc: "03", input: "abbcegjk", expected: false},
		{desc: "04", input: "abcdffaa", expected: true},
		{desc: "05", input: "ghjaabcc", expected: true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := validate11(tC.input)
			if actual != tC.expected {
				t.Fatalf("%s - Expected: %v got: %v\n", tC.input, tC.expected, actual)
			}
		})
	}
}

func TestNext11(t *testing.T) {
	t.Skip()
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{desc: "01", input: "aaaaaaaa", expected: "aaaaaaab"},
		{desc: "02", input: "aaaaaaaz", expected: "aaaaaaba"},
		{desc: "03", input: "azzzzzzz", expected: "baaaaaaa"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := next11(tC.input)
			if actual != tC.expected {
				t.Fatalf("%s - Expected: %v got: %v\n", tC.input, tC.expected, actual)
			}
		})
	}
}

func TestNextValid11(t *testing.T) {
	t.Skip()
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{desc: "01", input: "abcdefgh", expected: "abcdffaa"},
		{desc: "02", input: "ghijklmn", expected: "ghjaabcc"},
		{desc: "03", input: "vzbxkghb", expected: "vzbxxyzz"},
		{desc: "04", input: "vzbxxyzz", expected: "vzcaabcc"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := nextValid11(tC.input)
			if actual != tC.expected {
				t.Fatalf("%s - Expected: %v got: %v\n", tC.input, tC.expected, actual)
			}
		})
	}
}
