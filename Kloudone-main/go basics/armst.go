package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter an integer")
	fmt.Scanln(&n)
	temp := n
	var digit = 0

	for n > 0 {
		digit++
		n = n / 10
	}

	var sum int = 0
	n = temp
	var pow int
	for n > 0 {

		pow = 1
		for i := 0; i < digit; i++ {
			pow = pow * (n % 10)
		}

		sum = sum + pow
		n = n / 10
	}

	fmt.Println(temp)

	if sum == temp {
		fmt.Println(temp, "is an Armstrong number")
	} else {
		fmt.Println(temp, "is not Armstrong number")
	}

}
