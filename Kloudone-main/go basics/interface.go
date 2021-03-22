package main

import "fmt"

// Creating an interface

type tank interface {
	// Methods
	Area() float64
	Volume() float64
}

type values struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m values) Area() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m values) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank
	t = values{10, 14}
	fmt.Println("Area of tank ", t.Area())
	fmt.Println("Volume of tank ", t.Volume())
}
