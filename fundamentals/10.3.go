package main

import "fmt"

func main() {
	names := []string{"Andres", "Jose"}
	fmt.Println(names)

	names = append(names, "Antonio")
	fmt.Println(names)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println(names[1:])
}
