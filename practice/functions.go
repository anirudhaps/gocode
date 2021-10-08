package main

import "fmt"

// In go language, name of variable comes before its type
func sum(n1 int, n2 int) int {
	return n1 + n2
}

func printValue(num int) {
	fmt.Println(num)
}

// returning multiple values
func getNegetive(num int) (int, bool) {
	if (num >= 0) && (num < 25) {
		return -num, true
	}
	return num, false
}

// returning three values
func retValues() (int, bool, bool) {
	return 3, true, false
}

// If all the parameters are of same type, then the following syntax also works
func average(n1, n2 float32) float32 {
	return (n1 + n2) / 2.0
}

func main() {
	total := sum(11, 34)
	printValue(total)

	val, isValid := getNegetive(10)
	if isValid {
		printValue(val)
	} else {
		fmt.Println("Invalid value:", val)
	}

	// just checking the value's validity, assigning the actual value to _
	// this is same like discarding
	_, isValid = getNegetive(41)
	if isValid {
		fmt.Println("Valid value")
	} else {
		fmt.Println("Invalid value")
	}

	// Note: assigning something to _ discards that value.
	// This is useful when some value(s) returned by a funtion is/are not useful
	_, _, third := retValues() // discard first two values
	fmt.Println("Got third val:", third)

	fmt.Println("average of 17, 22:", average(17, 22))
}
