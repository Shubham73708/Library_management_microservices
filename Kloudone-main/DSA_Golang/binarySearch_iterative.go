package main

import "fmt"

func check(arr []int, x int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := l + (r-l)/2

		if arr[m] == x {
			return m
		}

		if arr[m] > x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
func main() {
	arr := []int{2, 3, 4, 5}
	x := 4
	res := check(arr, x)
	if res == -1 {
		fmt.Println("not found element")
	} else {
		fmt.Println(res)
	}
}
