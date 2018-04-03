package main

import (
	"fmt"
)

func main() {
	//var colors map[int]string
	colors := make(map[int]string)

	colors[10] = "#ff0000"
	colors[11] = "#ffffff"
	delete(colors, 10)
	fmt.Println(colors)
}
