package main

import "fmt"

func main() {

	i := 1
	x := 5
	fmt.Println(x)
	if x <= 6 {
		for i < 3 {
			i += 2
		}
		fmt.Println(i)
	}
}
