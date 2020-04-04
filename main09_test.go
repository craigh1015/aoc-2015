package main

import (
	"sort"
	"testing"
)

func routeEqual(a, b []route) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func pathEqual(a, b []path) bool {
	sort.Sort(byLocations(a))
	sort.Sort(byLocations(b))
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !stringEqual(v.locations, b[i].locations) || v.distance != b[i].distance {
			return false
		}
	}
	return true
}

func stringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestParseInput09(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected route
	}{
		{desc: "01", input: "Faerun to Norrath = 129", expected: route{from: "Faerun", to: "Norrath", distance: 129}},
		{desc: "02", input: "Tristram to AlphaCentauri = 118", expected: route{from: "Tristram", to: "AlphaCentauri", distance: 118}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := parseInput09(tC.input)
			if err != nil {
				t.Fatalf("Error parsing: %s - %v", tC.input, err)
			}
			if actual != tC.expected {
				t.Errorf("expected %v got %v", tC.expected, actual)
			}
		})
	}
}

func TestReverseRoute09(t *testing.T) {
	testCases := []struct {
		desc     string
		input    route
		expected route
	}{
		{desc: "01", input: route{from: "Faerun", to: "Norrath", distance: 129}, expected: route{to: "Faerun", from: "Norrath", distance: 129}},
		// {desc: "02", input: "Tristram to AlphaCentauri = 118", expected: route{from: "Tristram", to: "AlphaCentauri", distance: 118}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := reverseRoute09(tC.input)
			if actual != tC.expected {
				t.Errorf("expected %v got %v", tC.expected, actual)
			}
		})
	}
}

func TestLocations09(t *testing.T) {
	testCases := []struct {
		desc      string
		locations []route
		lookup    string
		expected  []route
	}{
		{desc: "01", locations: []route{}, lookup: "", expected: []route{}},
		{desc: "02", locations: []route{route{from: "Faerun", to: "Norrath", distance: 129}}, lookup: "Invalid name", expected: []route{}},
		{
			desc:      "03",
			locations: []route{route{from: "Faerun", to: "Norrath", distance: 129}, route{from: "Tristram", to: "AlphaCentauri", distance: 118}},
			lookup:    "Faerun",
			expected:  []route{route{from: "Faerun", to: "Norrath", distance: 129}},
		},
		{
			desc:      "04",
			locations: []route{route{from: "Faerun", to: "Norrath", distance: 129}, route{from: "Tristram", to: "AlphaCentauri", distance: 118}},
			lookup:    "Tristram",
			expected:  []route{route{from: "Tristram", to: "AlphaCentauri", distance: 118}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := lookupLocation09(tC.locations, tC.lookup)
			if !routeEqual(actual, tC.expected) {
				t.Errorf("expected %v got %v", tC.expected, actual)
			}
		})
	}
}

func TestProcess09(t *testing.T) {
	testCases := []struct {
		desc       string
		locations  []route
		poppedPath path
		expected   []path
	}{
		{
			desc:       "01",
			locations:  []route{route{from: "Faerun", to: "Norrath", distance: 129}, route{from: "Norrath", to: "AlphaCentauri", distance: 118}},
			poppedPath: path{locations: []string{"Faerun", "Norrath"}, distance: 129},
			expected:   []path{path{locations: []string{"Faerun", "Norrath", "AlphaCentauri"}, distance: 247}},
		},
		{
			desc: "02",
			locations: []route{
				route{from: "Faerun", to: "Norrath", distance: 129},
				route{from: "Norrath", to: "AlphaCentauri", distance: 118},
				route{from: "Norrath", to: "Tristram", distance: 121},
			},
			poppedPath: path{locations: []string{"Faerun", "Norrath"}, distance: 129},
			expected: []path{
				path{locations: []string{"Faerun", "Norrath", "AlphaCentauri"}, distance: 247},
				path{locations: []string{"Faerun", "Norrath", "Tristram"}, distance: 250},
			},
		},
		{
			desc: "03",
			locations: []route{
				route{from: "Faerun", to: "Norrath", distance: 129},
				route{from: "Norrath", to: "AlphaCentauri", distance: 118},
				route{from: "Norrath", to: "Faerun", distance: 129},
			},
			poppedPath: path{locations: []string{"Faerun", "Norrath"}, distance: 129},
			expected: []path{
				path{locations: []string{"Faerun", "Norrath", "AlphaCentauri"}, distance: 247},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := process09(tC.locations, tC.poppedPath)
			if !pathEqual(actual, tC.expected) {
				t.Errorf("expected %v got %v", tC.expected, actual)
			}
		})
	}
}

func TestGetPaths09(t *testing.T) {
	testCases := []struct {
		desc       string
		locations  []route
		poppedPath path
		expected   []path
	}{
		{
			desc: "01",
			locations: []route{
				route{from: "London", to: "Dublin", distance: 464},
				route{from: "London", to: "Belfast", distance: 518},
				route{from: "Dublin", to: "Belfast", distance: 141},
			},
			expected: []path{
				path{locations: []string{"Dublin", "London", "Belfast"}, distance: 982},
				path{locations: []string{"London", "Dublin", "Belfast"}, distance: 605},
				path{locations: []string{"London", "Belfast", "Dublin"}, distance: 659},
				path{locations: []string{"Dublin", "Belfast", "London"}, distance: 659},
				path{locations: []string{"Belfast", "Dublin", "London"}, distance: 605},
				path{locations: []string{"Belfast", "London", "Dublin"}, distance: 982},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := getPaths09(tC.locations)
			if !pathEqual(actual, tC.expected) {
				t.Errorf("expected %v got %v", tC.expected, actual)
			}
		})
	}
}

func TestGetStartPaths09(t *testing.T) {
	testCases := []struct {
		desc       string
		locations  []route
		poppedPath path
		expected   []path
	}{
		{
			desc: "01",
			locations: []route{
				route{from: "London", to: "Dublin", distance: 464},
				route{from: "London", to: "Belfast", distance: 518},
				route{from: "Dublin", to: "Belfast", distance: 141},
			},
			expected: []path{
				path{locations: []string{"London", "Dublin"}, distance: 464},
				path{locations: []string{"London", "Belfast"}, distance: 518},
				path{locations: []string{"Dublin", "Belfast"}, distance: 141},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := getStartPaths09(tC.locations)
			if !pathEqual(actual, tC.expected) {
				t.Errorf("expected %v got %v", tC.expected, actual)
			}
		})
	}
}
