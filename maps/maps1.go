package main

import (
	"fmt"
)

func main() {
	// This is an empty map, elements can be added to it
	colors := make(map[int]string)

	colors[10] = "#ff0000"
	colors[11] = "#ffffff"
	delete(colors, 10) // delete a key-value pair
	fmt.Println(colors)
	fmt.Println("colors[10]:", colors[10]) // empty string
	fmt.Println("len of map colors:", len(colors))

	// This is a nil map, so elements can't be added to it
	var student_age map[uint16]uint16
	// but we can try to access an element that is anyway non-existent
	// keys that are not present in a map have zero value (of the type) as corresponding value
	fmt.Println("student_age[5]:", student_age[5])
	// student_age[5] = 10 // error: assignment to entry in nil map
}
