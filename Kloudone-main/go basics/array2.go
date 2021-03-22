package main

import "fmt"

func main() {

	arr1 := [4]string{"ram", "namam", "vishu", "vinit"}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	arr := [3][3]int{{10, 5, 15},
		{0, 1, 6},
		{7, 2, 9}}

	fmt.Println("Elements of Array ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(arr[i][j])
		}
	}

}
