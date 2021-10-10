package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("First")
	fmt.Println("Second")

	time.Sleep(time.Second)
}
