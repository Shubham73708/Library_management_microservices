package main

import (
	"fmt"
	"time"
)

func Aname() {
	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}
	for t1 := 0; t1 <= 3; t1++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(arr1[t1])
	}
}

func count(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {

	go count("square")
	go count("cube")

	go Aname()
	time.Sleep(3500 * time.Millisecond)
	fmt.Println("...Main Go-routine End...")
}
