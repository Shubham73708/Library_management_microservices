// Golang program to illustrate the
// concept of encoding using JSON
package main

import (
	"encoding/json"
	"fmt"
)

// declaring a struct
type Human struct {

	// defining struct variables
	Name    string
	Age     int
	Address string
}

// main function
func main() {

	human1 := Human{"Ankit", 23, "New Delhi"}

	human_enc, err := json.Marshal(human1)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(string(human_enc))

	human2 := []Human{
		{Name: "Rahul", Age: 23, Address: "New Delhi"},
		{Name: "Pryank", Age: 20, Address: "Bulandshahr"},
		{Name: "Nihal", Age: 24, Address: "Noida"},
	}

	human2_enc, err := json.Marshal(human2)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println(string(human2_enc))
}
