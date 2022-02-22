package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("arf!")
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("meow!")
}

func speak(s Speaker) {
	s.Speak()
}

func main() {
	speak(&Dog{})
	speak(&Cat{})
}
