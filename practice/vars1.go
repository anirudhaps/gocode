package main

import "fmt"

func main() {
	name := "aps"
	// age, name := 30, 20
	// the above throws error: cannot use 20 (type untyped int) as type string in assignment
	age, name := 30, "aps"


	//fmt.Printf("%s's age is %d\n", name, age)
	fmt.Println("data:", name, age)
}