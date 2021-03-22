package main

import "fmt"

func main() {
	x := "124"
	arr := []rune(x)
	per(arr, 0, len(arr)-1)
}
func per(arr []rune, l int, r int) {
	if l == r {
		fmt.Println(string(arr))
	} else {
		for i := l; i <= r; i++ {
			arr[i], arr[l] = arr[l], arr[i]
			fmt.Println("Hi ", i, l)
			per(arr, l+1, r)
			arr[l], arr[i] = arr[i], arr[l]
		}
	}

}
