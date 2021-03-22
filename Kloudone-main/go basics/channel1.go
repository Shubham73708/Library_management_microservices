package main

import (
	"fmt"
	"time"
)

func myfun(ch chan int) {
	fmt.Println(234 + <-ch)
	time.Sleep(10 * time.Second)
	close(ch)
}
func main() {
	ch := make(chan int)
	go myfun(ch)
	ch <- 23
	fmt.Println("End Main method")
}
