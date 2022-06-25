package main

import "fmt"

func main() {
	input := []int{23, 1, 34, 56, 6, 7, 78, 9, 12, 23, 19, 20, 37, 75, 86}
	ch := make(chan int)
	for i := 0; i < len(input); i += 5 {
		go func(c chan int, inp []int) {
			fmt.Println("Goroutine:", inp)
			var sum int = 0
			for k := 0; k < len(inp); k++ {
				sum += inp[k]
			}
			c <- sum
		}(ch, input[i:i+5])
	}

	var total int = 0
	for j := 0; j < len(input); j += 5 {
		got := <-ch
		total += got
		fmt.Printf("got %d\n", got)
	}
	fmt.Printf("Total: %d\n", total)
}
