package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go func(c chan string, val int) {
			c <- "child goroutine: " + strconv.Itoa(val)
		}(ch, i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Main:", <-ch)
	}
}
