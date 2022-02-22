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
	fmt.Println(isAdult("male", 5))
	fmt.Println(isAdult("female", 19))
	fmt.Println(isAdult("alien", 10000))
}
