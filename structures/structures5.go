package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	aps := person{"Anirudha", "Singh"}
	fmt.Println(aps)
	ajay := person{firstName: "Ajay", lastName: "Anto"}
	fmt.Println(ajay)
	var amit person
	amit.firstName = "Amit"
	amit.lastName = "Tondon"
	fmt.Println(amit)
	fmt.Printf("%+v\n", amit)

	dummy := person{}
	fmt.Println(dummy.firstName, dummy.lastName)
}
