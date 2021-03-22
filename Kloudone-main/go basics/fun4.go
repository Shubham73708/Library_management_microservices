package main

import "fmt"

func main() {

	// Assigning an anonymous
	// function to a variable
	value := func() {
		fmt.Println("Welcome! in this industry")
	}
	value()

	func(ele string) {
		fmt.Println(ele)
	}("hello")
}
