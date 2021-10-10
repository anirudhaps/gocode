package main

import (
	"fmt"
	"time"
)

func StrPrint(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	StrPrint("First")

	go StrPrint("goroutine1")

	go func(msg string) {
		fmt.Println(msg)
	}("goroutine2")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
