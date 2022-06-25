package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(c chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println("Child goroutine:", i)
			c <- i
		}
	}(ch)

	for j := 0; j < 10; j++ {
		fmt.Println("Main:", <-ch)
	}
}
