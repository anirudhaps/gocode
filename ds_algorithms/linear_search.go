package main

import "fmt"

// data will receive a slice by pointer. This is default behaviour.
func search(target int, data []int) (int, bool) {
	for i, value := range data {
		if value == target {
			return i, true
		}
	}
	return -1, false
}

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	var target int

	fmt.Print("Enter a number to be searched in the list of primes below 30: ")
	fmt.Scanf("%d", &target)
	index, found := search(target, primes)
	if found {
		fmt.Printf("%d found at pos-[%d] in slice: %v\n", target, index, primes)
	} else {
		fmt.Printf("%d not found in slice: %v\n", target, primes)
	}
}
