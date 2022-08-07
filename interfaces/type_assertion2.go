package main

import "fmt"

type Operation interface {
	Add() int
}

type Numbers struct {
	x int
	y int
}

// Because Numbers is implementing Add(), it is implementing Operation interface
func (n *Numbers) Add() int {
	return n.x + n.y
}

func (n *Numbers) Subtract() int {
	return n.x - n.y
}

func main() {
	// assigning Numbers obj to Operation interface
	var oper Operation = &Numbers{20, 10}
	fmt.Printf("Sum: %v\n", oper.Add())
	// The below call to Subtract via oper will fail with error:
	// oper.Subtract undefined (type Operation has no field or method Subtract)
	// fmt.Printf("Diff: %v\n", oper.Subtract())
	// This means we can only invoke the methods declared in the interface via
	// that interface's variable even though the underlying struct (that is
	// implementing that interface) has that method. Thus, in order to invoke
	// other methods of Numbers, we have to type cast oper to *Numbers. Type
	// casting an interface to underlying type us known as type assertion.
	if num, ok := oper.(*Numbers); ok {
		fmt.Printf("Diff: %v\n", num.Subtract())
	}
}
