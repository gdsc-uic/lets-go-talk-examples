package main

import (
	"io/ioutil"
)

func main() {
	contents := []byte("hello peeps!")
	err := ioutil.WriteFile("./result.txt", contents, 0644)
	if err != nil {
		panic(err)
	}
}
