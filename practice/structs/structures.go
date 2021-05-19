package main

import "fmt"

// structure declaration
// name, age are fields
type Student struct {
	name string
	age int
}

func printInfo(name string, age int) {
	fmt.Println(name, age)
}

func main() {
	// ,(comma) after assigning value to a field is important
	ajit := Student{
		name: "Ajit",
		age: 28,
	}
	printInfo(ajit.name, ajit.age)

	// skip assigning both the values in this definition
	vivek := Student{}
	printInfo(vivek.name, vivek.age)

	// skip assigning one of the value during struct definition and then later
	// assign it using = operator
	jagadish := Student{name: "Jagadish"}
	jagadish.age = 19
	printInfo(jagadish.name, jagadish.age)

	// assigning both the values without field name. Here, order of values are
	// important
	suman := Student{"Suman", 25}
	printInfo(suman.name, suman.age)
}