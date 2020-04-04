package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func main12() {
	f, err := os.Open("input-12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	total := 0.0

	s := bufio.NewScanner(f)
	for s.Scan() {
		total += getNumbersTotalJSON12(s.Text())
	}

	log.Printf("Total: %.0f\n", total)
}

func getNumbersTotalJSON12(input string) float64 {
	var f interface{}
	err := json.Unmarshal([]byte(input), &f)
	if err != nil {
		log.Fatalf("Unable to parse %s - %v", input, err)
	}

	return processElement12(f)
}

func processElement12(element interface{}) float64 {
	switch m := element.(type) {
	case []interface{}:
		total := 0.0
		for _, v := range m {
			total += processElement12(v)
		}
		return total
	case map[string]interface{}:
		total := 0.0
		for _, v := range m {
			if v == "red" {
				return 0
			}
			total += processElement12(v)
		}
		return total
	case float64:
		return m
	default:
		return 0
	}
}
