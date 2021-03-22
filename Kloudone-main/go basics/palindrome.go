package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter any integer")
	fmt.Scanln(&n)
	temp := n
	var sum int = 0
	for n > 0 {
		sum = sum*10 + n%10
		n = n / 10
	}
	if temp == sum {
		fmt.Println(temp, "is Palindrome")
	} else {
		fmt.Println(temp, "is not Palindrome")
	}
}
