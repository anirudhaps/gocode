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
	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v\n", amit)

	dummy := person{}
	fmt.Println(dummy.firstName, dummy.lastName)

	// comma is important here after lastName line
	hiren := person{
		firstName: "Hiren",
		lastName:  "Shah",
	}
	fmt.Println(hiren.firstName, hiren.lastName)
	fmt.Printf("%v\n", hiren)
}
