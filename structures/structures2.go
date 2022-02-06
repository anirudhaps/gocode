package main

import "fmt"

type Student struct {
	name   string
	rollNo int
}

// associating functions to structure (like making them its member functions)
func (s *Student) display() {
	fmt.Printf("name: %s, roll no: %d\n", s.name, s.rollNo)
}

func (s *Student) updateRollNo() {
	s.rollNo++
}

// The below two method are acc. to factory pattern
// factory method for constructing struct objects
func NewStudent(sname string, sroll int) Student {
	return Student{
		name:   sname,
		rollNo: sroll,
	}
}

// another factory that returns pointer to a Student
func CreateStudent(sname string, sroll int) *Student {
	return &Student{
		name:   sname,
		rollNo: sroll,
	}
}

func main() {
	ajit := &Student{
		name:   "Ajit",
		rollNo: 1,
	}

	ajit.display()
	ajit.updateRollNo()
	ajit.display()
	fmt.Println("ajit's roll number:", ajit.rollNo)

	ajay := NewStudent("Ajay", 12)
	ajay.display()

	virat := CreateStudent("Virat Kohli", 33)
	virat.display()

	// using new function to allocate memory for a type like struct
	mallesh := new(Student)
	/*
		Above code is same as:
		mallesh := &Student{}
	*/
	mallesh.name = "Mallesham"
	mallesh.rollNo = 31

	mallesh.display()
}
