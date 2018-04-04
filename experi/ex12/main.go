package main

import (
    "fmt"
)

type point struct {
    x int
    y int
}

func main() {
    p1 := point{1, 2}
    p2 := point{}

    fmt.Println(p1)
    fmt.Println(p2)

    ptr := &p2
    (*ptr).x = 5
    fmt.Println(p2)
    ptr.y = 6
    fmt.Println(p2)
    p1.x = 9
    fmt.Println(p1)
    structStudy()
    arrayStudy()
    sliceStudy()
    sliceStudy1()
}

func structStudy() {
    var (
        p1 = point{2, 3}
        p2 = point{}
        p3 = point{y: 4, x: 9}
        p4 = point{x: 5}
        ptr = &point{10, 19}
    )
    fmt.Println("p1: ", p1)
    fmt.Println("p2: ", p2)
    fmt.Println("p3: ", p3)
    fmt.Println("p4: ", p4)
    fmt.Println("*ptr: ", *ptr)
    fmt.Println("ptr: ", ptr)
}

func arrayStudy() {
    var arr [2]string
    primes := [5]int {2, 3, 5, 7, 11}

    arr[0] = "Anirudha"
    arr[1] = "Singh"
    fmt.Println(arr)
    fmt.Println(arr[0], arr[1])
    fmt.Println(primes)

    for _, val := range primes {
        fmt.Println(val)
    }
}

func sliceStudy() {
    nums := [7]int{1, 2, 3, 5, 7, 9, 13}
    fmt.Println(nums)

    a := nums[0:3]
    b := nums[1:]

    fmt.Println("a:", a)
    fmt.Println("b:", b)
    a[1] = 19
    fmt.Println("a:", a)
    fmt.Println("b:", b)
    fmt.Println(nums)
    fmt.Println("nums[:]", nums[:])
}

func sliceStudy1() {
    nums := []int{1, 2, 3}
    bvals := []bool{true, false, true}
    fmt.Println("nums:", nums)
    fmt.Println("bvals:", bvals)

    svals := []struct{
        n int
        b bool
    }{
        {1, true},
        {2, false},
        {3, false},
    }
    fmt.Println("struct:", svals)
}
