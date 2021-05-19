package main

import "fmt"

func Maps() {
	// var students map[string]int
	// The above declaration creates a nil map
	// The default value for map is nil

	// The following assignments will result in a runtime error:
	// panic: assignment to entry in nil map

	//students["Ajit"] = 12
	//students["Vijay"] = 13
	//fmt.Println(students)
}

func Maps1() {
	students := make(map[string]int) // empty map at this point
	fmt.Println(students)            // map[]
	students["Ajit"] = 12
	students["Vijay"] = 13
	fmt.Printf("students: %v\n", students) // map[Ajit:12 Vijay:13]
	fmt.Printf("len: %d\n\n", len(students))
}

func Maps2() {
	students := make(map[string]int, 100) // empty map, preallocate room for 100 entries
	fmt.Println(students)                 // map[]
	fmt.Printf("len: %d\n\n", len(students))
}

func Maps3() {
	// another way of map declaration and initialization
	students := map[string]int{
		"Abhinav": 10,
		"Rakesh":  19,
	}
	fmt.Println(students)
	// len() returns number of key-value pairs
	fmt.Printf("len: %d\n\n", len(students))
}

func Maps4() {
	students := map[string]int{
		"Shah":    1,
		"Abdul":   2,
		"Xiang":   3,
		"Johnson": 4,
	}
	fmt.Printf("stud: %v, len: %d\n", students, len(students))
	roll := students["Xiang"]
	fmt.Printf("roll num of Xiang: %d\n", roll)

	roll, found := students["Xiang"]
	// when we index a map, it returns two values. The second one is a bool
	// that tells if the key exists in the map.
	if found {
		fmt.Printf("Xiang studies in this class at roll number: %d\n", roll)
	}

	// Trying with unknown key
	roll, found = students["Jai"]
	fmt.Printf("Is jai in class?: %t, so the roll number: %d\n", found, roll) // false, 0
	// In case of key not found in the map, the value returned is zero value (0 in above case)

	// just checking if a key exists in the map, discarding value:
	_, found = students["Abdul"]
	fmt.Printf("Is Abdul in class? %t\n", found)
	_, found = students["Ajay"]
	fmt.Printf("Is Ajay in class? %t\n", found)

	// delete a key-value pair
	delete(students, "Shah")
	fmt.Printf("stud: %v, len: %d\n", students, len(students))

	// trying to delete an unknown key
	delete(students, "Firoz") // silently ignores the key
	// the above way for deletion of non-existent key is not good
	// better to first see if the key is found in the map and based on that
	// delete it.

	// using for-each range loop for printing the map
	for key, value := range students {
		fmt.Printf("[%s]: %d\n", key, value)
	}
}

func main() {
	Maps()
	Maps1()
	Maps2()
	Maps3()
	Maps4()
}
