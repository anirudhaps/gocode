package main

import "fmt"

func main() {
	sum, avg := getSumAvg(1, 2, 3)
	fmt.Println("Sum of nums:", sum)
	fmt.Println("Avg of nums:", avg)
}

func getSumAvg(num1 int, num2 int, num3 int) (int, float64) {
	sumOfnums := num1 + num2 + num3
	avg := float64(sumOfnums) / 3
	return sumOfnums, avg
}
