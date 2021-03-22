package main

import "fmt"

func main() {

	func(s string) {
		fmt.Println(s)
	}("Welcome to the world")
}
