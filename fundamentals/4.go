package main

import "fmt"

func main() {
	a := 6
	if a%2 == 0 {
		fmt.Printf("%d is an even!\n", a)
	} else if a%3 == 0 {
		fmt.Printf("%d is divisible by 3!", a)
	} else {
		fmt.Printf("%d is not even!\n", a)
	}
}
