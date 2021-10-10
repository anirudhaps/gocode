package main

import "fmt"

type Student struct {
	name string
	age  int16
}

func (s Student) getName() string {
	return s.name
}

func (s *Student) getAge() int16 {
	return s.age
}

func main() {
	amit := &Student{
		name: "Amit",
		age:  30,
	}
	fmt.Println("name:", amit.getName())
	fmt.Println("age:", amit.getAge())
}
