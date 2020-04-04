package main

import "log"

func main11() {
}

func nextValid11(pwd string) string {
	result := pwd
	for {
		result = next11(result)
		if validate11(result) || result == "zzzzzzzz" {
			log.Printf("%s -> %s\n", pwd, result)
			return result
		}
	}
}

func validate11(pwd string) bool {
	return validateStraight11(pwd) && validateLetters11(pwd) && validatePairs11(pwd)
}

func next11(pwd string) string {
	bytes := []byte(pwd)
	for i := len(pwd) - 1; i >= 0; i-- {
		if bytes[i] != 'z' {
			bytes[i]++
			break
		}
		bytes[i] = 'a'
	}
	return string(bytes)
}

func validateStraight11(pwd string) bool {
	count := 0
	current := byte(0)
	for _, b := range []byte(pwd) {
		if current == b-1 {
			count++
		} else {
			count = 1
		}
		if count >= 3 {
			return true
		}
		current = b
	}
	return false
}

func validateLetters11(pwd string) bool {
	for _, b := range []byte(pwd) {
		if b == byte('i') || b == byte('o') || b == byte('l') {
			return false
		}
	}
	return true
}

func validatePairs11(pwd string) bool {
	var pair1 byte
	var pair1pos int
	for i := 0; i < len(pwd)-1; i++ {
		if pwd[i] == pwd[i+1] {
			pair1 = pwd[i]
			pair1pos = i
			break
		}
	}

	if pair1 == 0 {
		return false
	}

	for i := pair1pos + 2; i < len(pwd)-1; i++ {
		if pwd[i] != pair1 && pwd[i] == pwd[i+1] {
			return true
		}
	}

	return false
}
