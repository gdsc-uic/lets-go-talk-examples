package main

import "fmt"

func main() {
	color := "red"

	switch color {
	case "red", "blue", "yellow":
		fmt.Printf("%s is a primary color\n", color)
	case "black":
		fmt.Println("this is black")
	default:
		fmt.Println("not a color!")
	}
}
