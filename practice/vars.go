package main

import "fmt"

func main() {
	var a = 1  // declaration with init value

	// multiple variables declaration along with explicit type
	var b, c int = 5, 6
	d := 7
	name := "Anirudh"  // equivalent to: var name = "Anirudh"
	var friend = "Mahesh"

	// variable declaration without init value requires type to be mentioned in the declaration
	// Otherwise, compile time error
	var n int

	// multiple variable declaration in same line with initialization
	var bro, sis = "raj", "himani"
	var some_guy string
	// note: a var declaration without init value will set zero-value in the variable based on type
	// for eg. zero-value of type int is 0, for type string is "", for float is 0.0

	var some_float float64 = 2.345

	fmt.Println(a) 
	fmt.Println(b, c) 
	fmt.Println(d) 
	fmt.Println(name) 
	fmt.Println(n) 
	fmt.Println(friend) 
	fmt.Println(bro, sis) 
	fmt.Println(some_guy) 
	fmt.Println(some_float) 
}