package main

import "fmt"

func main() {
	var value = 5

	switch value {
	case 1:
		fmt.Println("C#")
	case 2, 3:
		fmt.Println("Go")
	case 4, 5:
		fmt.Println("Java")
	}
}
