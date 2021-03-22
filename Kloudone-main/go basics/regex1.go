package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("Hello")
	str := "Hello, how are you"
	match := re.FindStringIndex(str)
	fmt.Println(match)

	re2, _ := regexp.Compile("[0-9]+-v.*g")
	match1 := re2.FindString("20024-vani_gupta")
	fmt.Println(match1)
}
