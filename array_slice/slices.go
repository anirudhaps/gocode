package main

import (
	"fmt"
)

// slice declaration: type-1
func Slices1() {
	rolls := make([]int, 5)
	fmt.Println("Capacity:", cap(rolls)) // 5
	fmt.Println("Length:", len(rolls)) // 5
}

// slice declaration: type-2
func Slices2() {
	rolls := make([]int, 3, 5)
	fmt.Println("Capacity:", cap(rolls)) // 5
	fmt.Println("Length:", len(rolls)) // 3
}

// slice declaration: error case (length > capacity)
/*func Slices3() {
	rolls := make([]int, 5, 3)
	// compile time error: len larger than cap in make([]int)
}*/

// slice declaration: type-3
// Based on the count of literals in the slice, capacity and len will be set
func Slices4() {
	rolls := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Capacity:", cap(rolls)) // 6
	fmt.Println("Length:", len(rolls)) // 6
}

// slice declaration: type-3 (a different technique: declaring with index position)
// assigning a value at index 19 will ensure that the slice has at least 20 as
// length and capacity
func Slices5() {
	rolls := []int{19: 200}
	fmt.Println("Capacity:", cap(rolls)) // 20
	fmt.Println("Length:", len(rolls)) // 20
}

/* Note: arr := [4]int{1, 2, 3, 4}  is an array declaration
 * slice := []int{1, 2, 3, 4} is slice declaration with capacity and length as 4
 */

// slice declaration: type-4 (nil slice)
func Slices6() {
	var slice []int16
	fmt.Printf("slice: %v\n", slice) // []
	fmt.Printf("slice == nil: %t\n", slice == nil) // true
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice)) // 0, 0
}

// slice declaration: type-5 (empty slice)
func Slices7() {
	sliceOne := []int{}
	fmt.Printf("slice: %v\n", sliceOne) // []
	fmt.Printf("slice == nil: %t\n", sliceOne == nil) // false
	fmt.Printf("len: %d, cap: %d\n", len(sliceOne), cap(sliceOne)) // 0, 0

	// Another way:
	//sliceTwo := make([]int, 0)
}

func Slices8() {
	sliceOne := []int{10, 20, 30, 40, 50}
	fmt.Printf("%v\n", sliceOne)
	sliceOne[1] = 70
	fmt.Printf("%v\n", sliceOne)
}

// Taking a slice of a slice
func Slices9() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Printf("%v\n", slice)
	sliceOne := slice[1:3]
	fmt.Printf("sliceOne: %v, slice: %v\n", sliceOne, slice)
	// sliceOne: [20, 30], slice: [10, 20, 30, 40, 50]
	fmt.Printf("len: %d, cap: %d\n", len(sliceOne), cap(sliceOne)) // len: 2, cap: 4

	sliceOne[1] = 70
	fmt.Printf("sliceOne: %v, slice: %v\n", sliceOne, slice)
	// sliceOne: [20, 70], slice: [10, 20, 70, 40, 50]
	// The above output shows that the slice in a slice is just a window in original slice

	// If original slice has len: 5 (olen) and cap: 5 (ocap), then formula for len and cap
	// of slice out of it:
	// nslice = oslice[i:j] => len(nslice) == (j-i), cap(nslice) == (ocap - i)

	// sliceOne[3] = 100
	// The above gives: panic: runtime error: index out of range [3] with length 2
	// Reason: len(sliceOne) is 2 (0, 1 indexes), so we can't access beyond the length
	// even though the capacity is higher (4)
}

func main() {
	Slices1()
	Slices2()
	//Slices3()
	Slices4()
	Slices5()
	Slices6()
	Slices7()
	Slices8()
	Slices9()
}