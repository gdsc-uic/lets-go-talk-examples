package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Speak() {
	fmt.Printf("Hi my name is %s and I am %d years old!\n", p.name, p.age)
}

func main() {
	bob := &Person{name: "Bob", age: 20}
	bob.Speak()
}
