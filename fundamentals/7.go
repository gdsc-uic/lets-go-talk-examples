package main

import "fmt"

func isAdult(gender string, age int) bool {
	if gender == "male" {
		if age >= 20 {
			return true
		} else {
			return false
		}
	} else if gender == "female" {
		if age >= 18 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	if isAdult("male", 5) {
		fmt.Println("first person is an adult!")
	} else {
		fmt.Println("first person is not an adult!")
	}

	if isAdult("female", 19) {
		fmt.Println("second person is an adult")
	} else {
		fmt.Println("second person is not an adult")
	}
}
