package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go f1(c1)
	go f2(c2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Message received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Message received:", msg2)
		}
	}
}

func f1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "one"
}

func f2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "two"
}
