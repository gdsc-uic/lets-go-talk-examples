package main

import "fmt"

type Person struct {
	name string
	age  int
	place string
}

type Employee struct {
	Person
	position string
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

	em := Employee{
		Person:   p,
		position: "programmer",
	}
	fmt.Println(em.name)
	fmt.Println(em.age)
	fmt.Println(em.position)
}
