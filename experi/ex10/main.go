package main

import "fmt"

func main() {
    fmt.Println(sumOf(12, 21))
    x, y := swap("Hi", "There")
    fmt.Println(x, y)
    a, b := split(9)
    fmt.Println(a, b)
    printVar()
    printVarTypes()
    typeConvert()
    constTypes()
    loopStudy()
    loopStudy1()
    loopStudy2()
    sprintStudy()
}

func loopStudy() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Printf("Sum is %v\n", sum)
    //fmt.Printf("Last i value: %v\n", i)
}

// for is the while loop in go
func loopStudy1() {
    i := 0
    sum := 0
    for i < 10 {
        sum += i
        i++
    }
    fmt.Printf("While loop sum: %v\n", sum)
}

func loopStudy2() {
    i := 0
    sum := 0
    for {
        if i >= 10 {
            break
        }
        sum += i
        i++
    }
    fmt.Printf("Infinite loop sum: %v\n", sum)
}

func sprintStudy() {
    i := 5
    j := 10
    //Sprint converts to string and returns
    str := fmt.Sprint(i, j)
    fmt.Println(str)
}

func swap(a, b string) (string, string) {
    return b, a
}

func sumOf(x, y int) int {
    return x + y
}

//naked return
func split(s int) (x, y int) {
    x = s * 4/9
    y = s - x
    return
}

//learning var
//:= cannot be used outside a function
var python = true
func printVar() {
    var i, j int = 1, 2
    var k = 3
    n := 10
    fmt.Println(i, python, j, k, n)
}

func printVarTypes() {
    var (
        age int = 28
        name string = "APS"
        married bool = false
    )
    fmt.Printf("Type: %T  Value: %v\n", age, age)
    fmt.Printf("Type: %T  Value: %q\n", name, name)
    fmt.Printf("Type: %T  Value: %q\n", married, married)
}

func typeConvert() {
    fl := 2.0345
    i := int(fl)
    fmt.Println(fl, i)
}

func constTypes() {
    const p1 = 12
    fmt.Println("const", p1)
    //p1 = p1 + 12
}
