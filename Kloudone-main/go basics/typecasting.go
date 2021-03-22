package main

import "fmt"

func main() {
	var x int = 17
	var y int = 5
	var res float64

	res = float64(x) / float64(y)

	fmt.Println(res)
	var result int = int(res)
	fmt.Println(result)
}
