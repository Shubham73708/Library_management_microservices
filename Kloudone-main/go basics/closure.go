package main

import "fmt"

func main() {
	num := 5

	cube := func() int {
		num = num * num * num
		return num

	}
	fmt.Println(cube())
}
