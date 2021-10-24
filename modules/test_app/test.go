package main

import (
	"fmt"

	"test_app/morestrings"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("This is a test")
	fmt.Println(randomdata.SillyName())
	fmt.Println("Reverse of anirudha is:", morestrings.ReverseRunes("anirudha"))
}
