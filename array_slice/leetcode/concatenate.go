package main

import "fmt"

// implementation approach-1
/*func getConcatenation(nums []int) []int {
	result := make([]int, 2*len(nums))
	result = append(nums, nums...)
	return result
}*/

func getConcatenation(nums []int) []int {
	result := []int{}
	for i := 0; i < 2; i++ {
		for _, n := range nums {
			result = append(result, n)
		}
	}
	return result
}

func main() {
	numbers := []int{1, 2, 1}
	fmt.Println("Result:", getConcatenation(numbers))
}
