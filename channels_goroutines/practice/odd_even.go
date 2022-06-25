package main

import (
	"fmt"
	"sync"
)

// Wait group links:
// https://gobyexample.com/waitgroups
// https://www.educative.io/edpresso/what-are-wait-groups-in-golang
// https://www.educative.io/edpresso/how-to-use-waitgroup-in-golang

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	input := []int{23, 1, 34, 56, 6, 7, 78, 9, 12, 23, 19, 20, 37, 75, 86}

	go func() {
		defer wg.Done()
		for i := 0; i < len(input); i++ {
			if input[i]%2 == 0 {
				fmt.Printf("Even: %d\n", input[i])
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < len(input); i++ {
			if input[i]%2 != 0 {
				fmt.Printf("Odd: %d\n", input[i])
			}
		}
	}()
	wg.Wait()
}
