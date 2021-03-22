package main

import "fmt"

func main() {
	score := make(map[string]int)
	fmt.Println(score)
	score["Ritesh"] = 24
	score["Austin"] = 10
	score["Rock"] = 12
	fmt.Println(score)
	delete(score, "Rock")
}