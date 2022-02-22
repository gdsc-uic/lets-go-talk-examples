package main

import "fmt"

func isEven(num int) (bool, error) {
	if num == 0 {
		return false, fmt.Errorf("number is zero")
	}

	if num%2 == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	isEven1, err := isEven(2)
	if isEven1 {
		fmt.Println("2 is even!")
	} else if err != nil {
		panic(err)
	}

	isEven2, err := isEven(0)
	if isEven2 {
		fmt.Println("0 is even!")
	} else if err != nil {
		panic(err)
	}
}
