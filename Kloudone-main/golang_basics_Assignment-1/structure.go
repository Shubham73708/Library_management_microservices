package main

import "fmt"

type Address struct {
	Name string
	City string
	zip  int
}

func main() {
	var a Address
	fmt.Println(a)

	v1 := Address{"UP", "Noida", 3623572}
	fmt.Println("Address1: ", v1)

	v2 := Address{Name: "Bihar", City: "Ballia",
		zip: 277001}
	fmt.Println("Address2: ", v2)

	v3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", v3)

	fmt.Println("City Name: ", v1.Name)
	fmt.Println("City zip: ", v1.zip)
}
