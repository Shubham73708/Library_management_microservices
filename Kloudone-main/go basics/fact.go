package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter any integer")
	fmt.Scanln(&n)
	fmt.Println("Factorial of ", n, " is ", fact(n))
}
func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fact(n-1)

}
