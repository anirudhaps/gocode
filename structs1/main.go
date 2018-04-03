package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	ajay := person{
		firstName: "Ajay",
		lastName:  "Anto",
		contact: contactInfo{
			email:   "ajay@gmail.com",
			zipCode: 58000,
		},
	}
	ajay.print()
	ajay.updateName("Vijay")
	ajay.print()
	//ptrAjay := &ajay
	//ptrAjay.updateName2("Vijay")
	ajay.updateName2("Vijay") //shortcut
	ajay.print()
}

func (p person) updateName(newName string) {
	p.firstName = newName
}

func (p *person) updateName2(newName string) {
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
