package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go printFiboNums(c1, c2)
	fibonacci(c1, c2)
}

func printFiboNums(ch chan int, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
	quit <- 0
}

func fibonacci(ch chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done printing all the fibonacci numbers")
			return
		}
	}
}
