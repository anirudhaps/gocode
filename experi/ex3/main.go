package main

import "fmt"

func main() {
	//declation and initialization of a slice aka. resizable array
	nums := []int {1, 2, 3}
	cards := []string {"Ace of Spades", "Five of diamonds", newCard()}
	for i, val := range nums {
		fmt.Println(i, val)
	}
	for i, val := range cards {
		fmt.Println(i, val)
	}
	cards = append(cards, "Ace of diamonds")
	for i, val := range cards {
		fmt.Println(i, val)
	}
	nums = append(nums, 4)
	for i, val := range nums {
		fmt.Println(val)
	}
}

func newCard() string {
	return "six of diamonds"
}
