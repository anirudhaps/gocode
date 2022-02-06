package main

import "fmt"

// A new type definition "Age" which is a sort of alias of type uint32
type Age uint32

// we can define methods for this new user-defined type
func (a Age) toString() string {
	return fmt.Sprintf("%d", a)
}

// This is not a new type definition, this is just an alias of uint32 which is a basic type
type SomeAge = uint32

// hence, we can not define a method for a basic type (also known as non-local type).
// The following attempt to create a method for a basic type generates the following compiler
// error: cannot define new methods on non-local type uint32
//func (m SomeAge) getString() string {
//	return fmt.Sprintf("%d", m)
//}

func main() {
	var myAge Age = 30
	fmt.Printf("My age: %s\n", myAge.toString())
}
