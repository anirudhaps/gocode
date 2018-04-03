package main

import "fmt"

func main() {
	greetings := "Hi, Anirudha"
	bgreetings := []byte(greetings)
	ngreetings := string(bgreetings)
	fmt.Println(greetings)
	fmt.Println(bgreetings)
	fmt.Println(ngreetings)
}
