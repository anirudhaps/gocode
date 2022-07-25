package main

import (
	"fmt"
	"time"
)

func waitForExpirationTime(waitTime time.Duration, stop chan int) bool {
	ticker := time.NewTicker(waitTime)
	defer ticker.Stop()

	select {
	case <-ticker.C:
		return true
	case <-stop:
		return false
	}
}

func main() {
	rCh := make(chan int)
	wCh := make(chan int)
	go func(waitTime time.Duration, rCh chan int, wCh chan int) {
		if expired := waitForExpirationTime(waitTime, rCh); expired {
			fmt.Println("Timer expired")
		} else {
			fmt.Println("Timer stopped")
		}
		wCh <- 1
	}(10*time.Second, rCh, wCh)

	//fmt.Println("Main waiting for 6s")
	time.Sleep(10 * time.Second)
	rCh <- 1
	fmt.Println("Main waiting...")
	<-wCh
}
