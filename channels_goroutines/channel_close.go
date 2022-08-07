package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(closeMe chan int) {
		close(closeMe)
	}(ch)
	val, ok := <-ch
	// Note: when channel is closed, the ok will be false and the channel will
	// return the zero value (0 here because of int type) for the channel's type
	if ok {
		fmt.Printf("Channel is open, value from channel: %v\n", val)
	} else {
		fmt.Printf("Channel is closed, value from channel: %v\n", val)
	}
	val, ok = <-ch
	fmt.Printf("val: %v, ok: %v\n", val, ok)
	// Reading from a closed channel will always yield ok: false and val: zero
	// value of the channel's type
}
