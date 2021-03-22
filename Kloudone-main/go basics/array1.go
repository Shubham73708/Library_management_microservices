package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5}

	fmt.Println("Length:", len(array))

	array = append(array, 6)
	fmt.Println("SLength after appernd:", len(array))

}
