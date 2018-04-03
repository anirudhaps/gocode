package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeUnix := time.Now().UnixNano()
	src := rand.NewSource(timeUnix)
	robj := rand.New(src)
	num := robj.Intn(100)
	fmt.Println("Random number:", num)
}
