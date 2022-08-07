package main

import (
	"context"
	"fmt"
	"sync"
)

// num <-chan int is a receive-only channel of type int
// num chan<- int is a send-only channel of type int

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	input := make(chan int)
	// num is a receive only channel denoted by <-chan declaration.
	// It means inside the below function, we can only receive values from it.
	// We can't send anything to it. Though, input is a two-way channel, when it
	// is passed to below function, inside this function, it becomes receive-only
	// channel. In the caller function, it remains a two-way channel.
	go func(ctx context.Context, num <-chan int) {
		defer wg.Done()
		result := 0
		for {
			select {
			case <-ctx.Done():
				// when cancel is called on ctx by parent goroutine, this Done() channel is closed
				fmt.Printf("Sum: %v\n", result)
				return
			case val := <-num:
				result += val
			}
		}
	}(ctx, input)
	for _, n := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("input: %v\n", n)
		input <- n
	}
	cancel()
	wg.Wait()
}
