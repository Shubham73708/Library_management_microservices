package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Something went wrong"

	match, err := regexp.MatchString("thing", str)
	fmt.Println("Match :", match)
	fmt.Println("Error :", err)

	match1, err := regexp.MatchString("          ", str)
	fmt.Println("match", match1, "Error", err)
}
