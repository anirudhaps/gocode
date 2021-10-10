package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, this is %s. I am %d years old.\n", p.name, p.age)
}

type Employee struct {
	//basicInfo *Person
	*Person   // composition of Person in Employee struct without giving field name
	department string
}
// Go will give name to Person that is "Person"

func (e *Employee) Introduce() {
	e.Person.Introduce()
	fmt.Printf("My department is %s\n", e.department)	
}

/* We can have one Introduce method of Person and
 * another Introduce method of Employee  
 */

func main() {
	mahesh := Employee{
		Person: &Person{
			name: "Mahesh Kurund",
			age: 30,
		},
		department: "SW",
	}
	//mahesh.basicInfo.Introduce()
	//mahesh.Introduce() // since we have not given any explicit field name, we can implicitly access the fields and functions of composed type
	//mahesh.Person.Introduce()
	//fmt.Printf("My department is %s\n", mahesh.department)
	mahesh.Introduce()
}