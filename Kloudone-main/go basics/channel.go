package main

import "fmt"

func mtain() {
	var cannel chan int
	fmt.Println("Value of channel: ", channel)
	fmt.Printf("Type of channel: %T ", channel)

	channel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", channel1)
	fmt.Printf("Type of the channel1: %T ", channel1)
}
