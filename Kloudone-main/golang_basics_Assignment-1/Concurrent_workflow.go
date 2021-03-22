package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var data chan string

func myfun() {
	fmt.Println("Goroutine fun is now starting...")
	defer func() {
		fmt.Println("Destroying the fun...")
		waitGroup.Done()
	}()

	for {

		value, ok := <-data
		if !ok {
			fmt.Println("The channel is closed!")
			break
		}

		fmt.Println(value)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Programm is starting...")
	data = make(chan string)

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go myfun()
	}

	for i := 0; i < 10; i++ {
		data <- ("Testing " + strconv.Itoa(i))
	}

	close(data)

	waitGroup.Wait()
}
