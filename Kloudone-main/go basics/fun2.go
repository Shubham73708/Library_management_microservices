package main

import "fmt"

func swap(a, b *int) {
	var temp = *a
	*a = *b
	*b = temp

}

func main() {
	var (
		a int
		b int
	)
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	fmt.Println(a, b)
	swap(&a, &b) // call by reference
	fmt.Println("After swap0", a, b)
}
