package main

import "fmt"

func main() {
	n := 0
	for n < 5 {
		fmt.Printf("n is %d\n", n)
		n++
	}

	for n := 20; n >= 10; n -= 2 {
		fmt.Println(n)
	}

	for {
		fmt.Println("this is looping")
	}
}
