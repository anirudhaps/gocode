package main

import "fmt"

func getName(roll int) string {
	var name string
	switch roll {
	case 1:
		name = "Ajay"
	case 2:
		name = "Vijay"
	case 3:
		name = "Nikita"
	default:
		name = "(unknown)"
	}
	return name
}

type Student struct {
	name string
	age  int8
}

// unlike c/c++ and other languages, non-constants are allowed in case expressions
func ageAnalysis(student *Student) {
	switch age := student.age; {
	case age < 2:
		fmt.Println("Infant")
	case age >= 2 && age < 3:
		fmt.Println("Toddler")
	case age >= 3 && age < 6:
		fmt.Println("Preschooler")
	case age >= 6 && age <= 12:
		fmt.Println("School-aged")
	case age >= 13 && age <= 18:
		fmt.Println("Teenager")
	default:
		fmt.Println("Adult")
	}
}

func getCharType(ch rune) string {
	var chType string
	switch ch {
	case '\n':
		fallthrough
	case '\t':
		fallthrough
	case ' ':
		fallthrough
	case '\r':
		chType = "whitespace"
	case 'a', 'e', 'i', 'o', 'u':
		chType = "Vowel"
	default:
		chType = "Consonent"
	}
	return chType
}

func typeChecker(val interface{}) {
	// this is called type switch
	switch val.(type) {
	case rune:
		fmt.Println(val, "is a rune")
	case string:
		fmt.Println(val, "is a string")
	default:
		fmt.Printf("Type of %v is %T\n", val, val)
	}
}

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(getName(i))
	}

	studs := []*Student{
		{"Ajay", 20},
		{"Vijay", 23},
		{"Jyoti", 5},
		{"Vivek", 10},
	}

	for i := 0; i < len(studs); i++ {
		ageAnalysis(studs[i])
	}

	fmt.Println("Analyse my name:")
	// range will return each character of string as rune
	for _, ch := range "Anirudha P Singh" {
		fmt.Printf("%s ", getCharType(ch))
	}
	fmt.Printf("\n")

	fmt.Println("Type analysis:")
	// each char in golang is a unicode character and thus is of type rune
	values := []interface{}{'j', "hello", 34, 2.34}
	for _, val := range values {
		typeChecker(val)
	}
}
