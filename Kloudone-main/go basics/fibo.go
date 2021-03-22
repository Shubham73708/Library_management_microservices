package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}
func main() {

	n := 9

	fmt.Println(fibo(n))
}
