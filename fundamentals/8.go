package main

import "fmt"

func main() {
	pp := struct {
		city string
	}{
		city: "Davao",
	}

	fmt.Println(pp.city)
}
