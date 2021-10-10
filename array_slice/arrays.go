package main

import (
	"fmt"
)

func main() {
	var names [4]string
	names[0] = "Amit"
	names[1] = "Vishal"
	names[2] = "Shiva"
	names[3] = "Shikha"

	// scope of i is limited to the loop
	for i := 0; i < len(names); i++ {
		fmt.Printf("%s", names[i])
		if i == len(names) - 1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" ")
		}
	}

	fmt.Printf("Using range:\n")
	for index, value := range names {
		fmt.Printf("[%d]: %s\n", index, value)
	}

	ages := [4]int{23, 21, 29, 25}

	// the following forms a while loop, here, scope of i is not limited to the loop
	i := 0
	for i < len(ages) {
		fmt.Printf("%d", ages[i])
		if i == len(ages)-1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" ")
		}
		i++
	}
	fmt.Println(ages)
}