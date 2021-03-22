package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func main() {
	var (
		a int
		b int
	)
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	fmt.Println("sum is", add(a, b))
	fmt.Println("subtraction is", sub(a, b))
	fmt.Println("product is", mul(a, b))
	fmt.Println("division is", div(a, b))
}
