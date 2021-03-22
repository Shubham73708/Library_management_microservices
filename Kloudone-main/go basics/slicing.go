package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Amit",
		"Jack",
		"Ritesh",
	}
	fmt.Println(names)
	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)
	slice2[0] = "Tom"
	fmt.Println(slice1, slice2)
	fmt.Println(names)
}
