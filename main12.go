package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main12() {
	f, err := os.Open("input-12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	total := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		total += getNumbersTotal12(s.Text())
	}

	log.Printf("Total: %d\n", total)
}

func getNumbersTotal12(input string) int {
	result := 0
	re := regexp.MustCompile(`[-\d]+`)
	numbers := re.FindAll([]byte(input), -1)
	for _, numberString := range numbers {
		numberValue, err := strconv.Atoi(string(numberString))
		if err != nil {
			log.Fatalf("Unable to parse %s - %v", numberString, err)
		}
		result += numberValue
	}
	return result
}
