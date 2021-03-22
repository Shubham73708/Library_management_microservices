package main

import "fmt"

func area(p, q int) (rectangle int, square int) {
	rectangle = p * q
	square = p * p
	return
}

func main() {
	var a1, a2 = area(2, 4)

	fmt.Printf("Rectangle area is: %d", a1)
	fmt.Println("\nSquare area is:", a2)
	fmt.Println(a1, a2)
}
