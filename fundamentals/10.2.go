package main

import "fmt"

func main() {
	names := []string{"Andres", "Jose"}
	fmt.Println(names)

	names = append(names, "Antonio")
	fmt.Println(names)
}
