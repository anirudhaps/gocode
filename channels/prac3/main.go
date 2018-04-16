package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go fibonacci(15, c)
	// for t := range c {
	// 	fmt.Print(t, " ")
	// }
	for t, ok := <-c; ok == true; {
		fmt.Print(t, " ")
		t, ok = <-c
	}
	fmt.Println()
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	ch <- x
	ch <- y
	for i := 0; i < (n - 2); i++ {
		ch <- x + y
		x, y = y, x+y
	}
	close(ch)
}
