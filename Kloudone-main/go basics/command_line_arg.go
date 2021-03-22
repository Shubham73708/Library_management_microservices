package main

import (
	"fmt"
	"os"
)

func main() {
	var cmd = os.Args[0]
	fmt.Println(cmd)
	cmd1 := os.Args[1]
	fmt.Println("second argument", cmd1)
}
