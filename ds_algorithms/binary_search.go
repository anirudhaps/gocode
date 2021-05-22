package main

import "fmt"

func search(target int, data []int) (int, bool) {
	var mid int
	beg := 0
	end := len(data) - 1

	for beg <= end {
		mid = (beg + end) / 2
		if data[mid] == target {
			return mid, true
		} else if target > data[mid] {
			beg = mid + 1
		} else {
			end = mid - 1
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
		fmt.Printf("Prime number %d found at pos-[%d]\n", target, index)
	} else {
		fmt.Printf("%d is not a prime number under 30\n", target)
	}
}
