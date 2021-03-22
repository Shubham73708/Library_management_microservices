package main

import "fmt"

func check(arr []int, l int, r int, x int) int {

	if l <= r {
		m := l + (r-l)/2

		if arr[m] == x {
			return m
		}

		if arr[m] > x {
			return check(arr, l, m-1, x)
		}

		return check(arr, m+1, r, x)

	}
	return -1

}
func main() {
	arr := []int{2, 3, 4, 5}
	x := 4
	res := check(arr, 0, len(arr)-1, x)
	if res == -1 {
		fmt.Println("not found element")
	} else {
		fmt.Println(res)
	}

}
