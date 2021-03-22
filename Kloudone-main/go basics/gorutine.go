package main

import "fmt"

func display(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Println(str)
	}
}
func main() {
	go display("Welcome")
	display("World")

}
