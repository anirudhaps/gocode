package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	fmt.Println("Starting main goroutine...")
	go printMsg("Hello", c)
	r := <-c
	if r == 1 {
		fmt.Println("Received val from child goroutine...")
	}
	fmt.Println("End of main goroutine.")
}

func printMsg(msg string, ch chan int) {
	fmt.Println(msg)
	ch <- 1
}
