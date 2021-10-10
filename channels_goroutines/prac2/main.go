package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	lst := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Starting main go routine...")
	go sum(lst[:len(lst)/2], c)
	go sum(lst[len(lst)/2:], c)
	s1 := <-c
	s2 := <-c
	fmt.Println("Received two sums from go routines:", s1, s2)
	fmt.Println("Total sum of list:", s1+s2)
}

func sum(list []int, ch chan int) {
	s := 0
	for i := 0; i < len(list); i++ {
		s += list[i]
	}
	ch <- s
}
