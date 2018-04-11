package main

import (
	"fmt"
)

type vertex struct {
	x int
	y int
}

var coords map[string]vertex

func main() {
	//coords := make(map[string]vertex)
	coords = make(map[string]vertex)
	fmt.Println(coords)
	//coords["longitude"] = vertex{
	//	23, 39,
	//}
	//coords["latitude"] = vertex{
	//45, 50,
	//}
	coords["longitude"] = vertex{23, 39}
	coords["latitude"] = vertex{38, 91}
	fmt.Println(coords)
	for i, val := range coords {
		fmt.Println("key:", i, "val:", val)
	}
}
