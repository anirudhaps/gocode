package main

import (
	"fmt"
)

// The interface{} type: the empty interface
// It has no methods.
// All types satisfy the empty interface.
// A function that has parameter of type interface{}, can receive value of any
// type
func PrintGeneric(val interface{}) {
	fmt.Println(val)
}

func main() {
	PrintGeneric("hello")
	PrintGeneric(24)
	PrintGeneric(50.567)
	PrintGeneric(true)
}
