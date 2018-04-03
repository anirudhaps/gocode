package main

import "fmt"

func main() {
	newCard()
	newi := getInt()
	fmt.Println(newi)
	newf := getFloat()
	fmt.Println(newf)
}

//This function des not return anything
func newCard() {
	fmt.Println("Ace of spades")
}

//This function returns an int
func getInt() int {
	return 25
}

//This function returns a float
func getFloat() float64 {
	return 3.14
}
