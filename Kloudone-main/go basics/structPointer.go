package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	emp := &Employee{"Sam", "Jon", 55, 15000}

	// (*emp8).firstName is the syntax to access
	// the firstName field of the emp8 struct
	fmt.Println("First Name:", (*emp).firstName)
	fmt.Println("Age:", emp.age)
}
