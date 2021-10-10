package main

import "fmt"

// structure declaration
// name, age are fields
type Student struct {
	name string
	age int
}

func displayInfo(student Student) {
	student.name = "Vijay"
	fmt.Printf("Student name: %s, Student age: %d\n", student.name, student.age)
}

func updateInfo(student *Student) {
	student.name = "Random"
	fmt.Printf("Student name: %s, Student age: %d\n", student.name, student.age)
}

func main() {
	ajit := Student{
		name: "Ajit",
		age: 28,
	}
	displayInfo(ajit)
	fmt.Println(ajit.name, ajit.age)

	// vivek stores address to the structure
	vivek := &Student{
		name: "Vivek",
		age: 27,
	}
	fmt.Println("Before update:", vivek.name, vivek.age)
	updateInfo(vivek)
	fmt.Println("After update:", vivek.name, vivek.age)
}