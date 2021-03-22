package main

import "fmt"

func main() {

	var a int = 10
	b := 25
	fmt.Println("before swap", a, b)
	swap(&a, &b)
	fmt.Println("after swap", a, b)
}

func swap(a, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}
