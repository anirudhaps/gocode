package main

import (
	"fmt"
)

func main() {
	val := "Bill"
	fmt.Println(val)
	updateValue(val)
	fmt.Println(val)
}

func updateValue(val string) {
	val = "Alex"
}
