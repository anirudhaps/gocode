package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures
type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

// There is no 'implement' keyword in golang.
// To implement an interface, we have to implement all the methods of the
// interface for a type

// Important note: The receiver can be pointer or a value for type-methods
// implementing the interface. The interface specification does not specify
// anything about receiver.
// implementing geometry interface for circle type
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// implementing geometry interface for rectangle type
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

// If the variable is of named interface type, then it can be used to call the
// methods declared in the interface.
// We can pass a variable of any type that implements geometry interface.
// Depending on the passed type var, the corresponding implemented methods will
// be invoked.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	rect := rectangle{width: 20, height: 3}
	cir := circle{radius: 3}
	measure(rect)
	measure(cir)

	// another way:
	// we can create a slice of type geometry and each element of that slice can
	// be a type that has implemented the interface geometry
	figures := []geometry{rectangle{10, 4}, circle{5}}
	//fmt.Println(figures)

	fmt.Println("calling measure() on each figure:")
	for _, val := range figures {
		measure(val)
	}
}
