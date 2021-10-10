package main

import (
	"fmt"
	"os"
)

func main() {
	// len() function can be used for getting the length of a string,
	// number of elements in an array or the number of values in a dictionary.
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("cmd args:", os.Args[0], os.Args[1])
}