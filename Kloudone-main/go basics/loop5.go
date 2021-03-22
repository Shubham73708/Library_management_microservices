package main

import "fmt"

func main() {

	for j := range "XabCd" {
		fmt.Printf("The index number is %d\n", j)
	}
}
