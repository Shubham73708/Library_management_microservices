package main

import (
	"errors"
	"fmt"
)

func main() {
	error := errors.New("There is an erro")
	fmt.Println(error)
	error2 := errors.New("This is no error")
	fmt.Println(error2)
}
