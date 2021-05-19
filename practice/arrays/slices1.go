package main

import (
	"fmt"
	"reflect"
)

func SliceAppend() {
	nums := []int{10, 20, 30, 40}
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))
	nums = append(nums, 50)
	fmt.Printf("nums: %v, len: %d, cap: %d\n", nums, len(nums), cap(nums))

	// append() returns a new slice after appending the value
	// append() increases the length by one. If capacity of the underlying
	// array is exhauted, append will increase the capacity by 2x
}

func SliceAppend1() {
	nums := []int{10, 20, 30, 40, 50, 60}
	slice := nums[1:3]
	fmt.Printf("nums: %v, slice: %v, len(nums): %d, cap(nums): %d, len(slice): %d, cap(slice): %d\n",
		nums, slice, len(nums), cap(nums), len(slice), cap(slice))

	slice = append(slice, 70)
	fmt.Printf("After append: nums: %v, slice: %v, len(nums): %d, cap(nums): %d, len(slice): %d, cap(slice): %d\n",
		nums, slice, len(nums), cap(nums), len(slice), cap(slice))

	newSlice := append(nums, 90)
	fmt.Printf("newSlice: %v, old(nums): %v\n", newSlice, nums)
	// changing this new slice will not change the old slice (nums) because the
	// newSlice has a new underlying array
}

func SliceAppend2() {
	slice1 := []int{10, 20}
	slice2 := []int{30, 40}

	slice3 := append(slice1, slice2...) // use triple dots (...) to append slice2 to slice1
	fmt.Println(slice3)
}

func SliceAppend3() {
	arr := [...]int{10, 20, 30, 40} // array (...) three dots create an array
	slice := []int{10, 20}          // slice
	fmt.Println("Type of arr:", reflect.TypeOf(arr))
	fmt.Println("Type of slice:", reflect.TypeOf(slice))
}

func SliceAppend4() {
	arr := [...]int{10, 20, 30, 40}
	s := arr[:]
	fmt.Printf("arr[:]: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	s = arr[1:]
	fmt.Printf("arr[1:]: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	s = arr[:2]
	fmt.Printf("arr[:2]: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	s = arr[1:3]
	fmt.Printf("arr[1:3]: %v, len: %d, cap: %d\n", s, len(s), cap(s))
}

func main() {
	SliceAppend()
	SliceAppend1()
	SliceAppend2()
	SliceAppend3()
	SliceAppend4()
}
