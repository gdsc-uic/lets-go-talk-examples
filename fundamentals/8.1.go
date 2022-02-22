package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	pp := struct {
		city string
	}{
		city: "Davao",
	}
	fmt.Println(pp.city)

	p := Person{name: "Ned", age: 18}
	fmt.Println(p.name)
	fmt.Println(p.age)
}
