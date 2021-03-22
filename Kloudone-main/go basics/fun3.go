package main

import (
	"fmt"
	"strings"
)

func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func main() {

	fmt.Println(joinstr())
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println(joinstr("Geeks", "for", "Geeks"))
	fmt.Println(joinstr("G", "E", "E", "k", "S"))

}
