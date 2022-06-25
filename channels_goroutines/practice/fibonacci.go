package main

import "fmt"

func Fibonacci(c chan int, q chan string) {
	var x, y int = 0, 1
	for {
		select {
		case c <- x:
			temp := x
			x = y
			y = temp + y
		case msg := <-q:
			if msg == "quit" {
				break
			}
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan string)
	go Fibonacci(ch, quit)

	var n int
	fmt.Printf("Enter the fibonacci number count: ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-ch)
	}
	fmt.Printf("\n")
	quit <- "quit"
}
