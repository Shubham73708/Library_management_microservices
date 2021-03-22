package main

import "fmt"

type student struct {
	rollNo int
	name   string
}

func main() {
	stud := student{rollNo: 1, name: "Mayank"}
	fmt.Println(stud.name, stud.rollNo)
	fmt.Println(student{rollNo: 2, name: "Rohan"})
}
