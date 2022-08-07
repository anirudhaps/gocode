package main

import (
	"fmt"
	"time"
)

func main() {
	lastSeen := time.Time{}
	fmt.Println("Zero Unix time:", uint32(lastSeen.Unix()))
}
