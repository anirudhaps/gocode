package main

import "fmt"

// nestead struct
type Employee struct {
	name string
	age int
	department string
	manager *Employee
}

func main() {
	mallesh := &Employee{
		name: "Mallesham",
		age: 31,
		department: "sw",
		manager: &Employee{
			name: "Bhyrav",
			age: 52,
			department: "management",
			manager: nil,
		},
	}

	fmt.Printf("Employee:\nname: %s\nage: %d\ndepartment: %s\nmanager:\n\tname: %s\n\tage: %d\n\tdepartment: %s\n\tmanager: %v\n",
		mallesh.name, mallesh.age, mallesh.department, mallesh.manager.name, mallesh.manager.age, mallesh.manager.department,
		mallesh.manager.manager)
}