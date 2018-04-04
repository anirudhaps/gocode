package main

import "fmt"

func main() {
    i := 25
    j := 20
    p := &i
    i++
    fmt.Printf("i = %v   *p = %v\n", i, *p)
    *p = 12
    fmt.Printf("i = %v   *p = %v\n", i, *p)
    p = &j
    *p = *p / 10
    fmt.Printf("i = %v   j = %v\n", i, j)
}
