package main

import (
	"fmt"
	"time"
)

func generateNumber(data chan int) {
	n := 0

	for {
		n++
		time.Sleep(1 * time.Second)
		if n < 10 {
			data <- n
		}
	}
}

func main() {
	dataChan := make(chan int)
	go generateNumber(dataChan)

	for {
		select {
		case n := <-dataChan:
			fmt.Printf("received number: %d\n", n)
		case <-time.After(1 * time.Second):
			fmt.Println("nothing to receive")
		}
	}
}
