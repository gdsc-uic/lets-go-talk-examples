package main

import (
	"fmt"
)

func main() {
	favoritePet := ""
	fmt.Print("input your pet: ")
	_, err := fmt.Scanln(&favoritePet)
	if err != nil {
		panic(err)
	}

	fmt.Printf("my favorite pet is %s\n", favoritePet)
}
