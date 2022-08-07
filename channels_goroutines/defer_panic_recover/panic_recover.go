package main

import "fmt"

func nonZero() (val int) {
	defer func() {
		switch p := recover(); p.(type) {
		case int:
			val = 5
		default:
			panic(p)
		}
	}()

	panic(1)
}

func main() {
	fmt.Println(nonZero())
}
