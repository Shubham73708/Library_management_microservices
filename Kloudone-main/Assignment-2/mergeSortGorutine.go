package main

import "fmt"

func main() {

	arr := []int{5, 2, 45, 50, 20}
	mergeSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func sortArray(nums []int) []int {
	l := len(nums)
	mergeSort(nums, 0, l-1)

	return nums
}

// merge sort
func mergeSort(nums []int, l int, r int) {
	if l == r {
		return
	}

	m := (l + r) / 2
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	go func() {
		mergeSort(nums, l, m)
		defer waitgroup.Done()
	}
	go func() {
		mergeSort(nums, m+1, r)
		defer waitgroup.Done()
	}

	waitgroup.Wait()
	merge(nums, l, m, r)
	//fmt.Println(nums)
}

func merge(nums []int, l int, m int, r int) {
	n1 := m - l + 1
	n2 := r - m
	L := make([]int, n1)
	for x := 0; x < n1; x++ {
		L[x] = nums[x+l]
	}

	R := make([]int, n2)
	for x := 0; x < n2; x++ {
		R[x] = nums[x+m+1]
	}

	a := l

	c := 0
	d := 0

	for {
		if (c > (n1 - 1)) || (d > (n2 - 1)) {
			break
		}

		if L[c] <= R[d] {
			nums[a] = L[c]
			c++

		} else {
			nums[a] = R[d]
			d++
		}
		a++
	}

	if c <= (n1 - 1) {
		for {
			if c > (n1 - 1) {
				break
			}

			nums[a] = L[c]
			c++
			a++
		}
	}

	if d <= (n2 - 1) {
		for {
			if d > (n2 - 1) {
				break
			}
			nums[a] = R[d]
			d++
			a++
		}
	}
	//fmt.Println(nums)
	return
}
