package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, this is %s. I am %d years old.\n", p.name, p.age)
}

type Employee struct {
	//basicInfo *Person
	// composition of Person in Employee struct without giving field name
	// More prominently called Embedding
	// Here Person is embedded in the Employee struct
	*Person
	department string
}

// Go will give name to Person that is "Person"

func (e *Employee) Introduce() {
	// e.Person.Introduce()
	// Embedding of Person into Employee struct ensures that name, age members
	// of the embedded Person struct are directly available in the Employee
	// struct without having to qualify them with Person.

	// Hence, this way is not needed
	// fmt.Printf("name: %s, age: %d\n", e.Person.name, e.Person.age)

	fmt.Printf("name: %s, age: %d\n", e.name, e.age)
	fmt.Printf("My department is %s\n", e.department)
}

/* We can have one Introduce method of Person and
 * another Introduce method of Employee
 */

func main() {
	mahesh := Employee{
		Person: &Person{
			name: "Mahesh Kurund",
			age:  30,
		},
		department: "SW",
	}
	//mahesh.basicInfo.Introduce()
	//mahesh.Introduce() // since we have not given any explicit field name, we can implicitly access the fields and functions of composed type
	//mahesh.Person.Introduce()
	//fmt.Printf("My department is %s\n", mahesh.department)
	mahesh.Introduce()
}
