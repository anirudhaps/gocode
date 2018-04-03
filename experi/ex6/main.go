package main

import "fmt"

func main() {
	names := []string {"Anirudha", "Ajay", "Nidhi", "Abhijit", "Rajendar", "Mallesham"}

	fmt.Println("name list:")
	printNames(names)

	sl1 := names[0:3]
	sl2 := names[3:]
	sl3 := names[:]
	sl4 := sl1
	sl5 := sl1[:]
	//sl6 := names[0:29]
	fmt.Println("list-1:")
	printNames(sl1)
	fmt.Println("list-2:")
	printNames(sl2)
	fmt.Println("list-3:")
	printNames(sl3)
	fmt.Println("list-4:")
	printNames(sl4)
	sl4[0] = "Vivek"
	fmt.Println(sl4[0], sl1[0])
	fmt.Println("list-5:")
	printNames(sl5)
	sl5[0] = "Harsha"
	fmt.Println(sl5[0], sl1[0])
	//fmt.Println("list-6:")
	//printNames(sl6)
}

func printNames(n []string) {
	for _, name := range n {
		fmt.Println(name)
	}
}
