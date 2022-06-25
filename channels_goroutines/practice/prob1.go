package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(c chan int) {
		fmt.Println("Child Goroutine:")
		c <- 1
	}(ch)
	r := <-ch
	fmt.Println("Value received by main goroutine from channel:", r)
	//fmt.Println("End of main")
}
