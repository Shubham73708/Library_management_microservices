package main

import "fmt"

//Bubble sort function
func bubblesort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

}

//main fumction
func main() {
	arr := []int{20, 13, 45, 40, 120, 42, 40}
	bubblesort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
