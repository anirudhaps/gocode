package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readType1() {
	var num int
	var name string

	fmt.Print("Enter any number and any string: ")
	count, err := fmt.Scanf("%d %s", &num, &name)
	fmt.Println("The number entered:", num)
	fmt.Println("The name entered:", name)
	fmt.Printf("count: %d, err: %v\n", count, err)
}

func readType2() {
	var num int
	var name string
	fmt.Print("Enter number and string: ")
	count, err := fmt.Scan(&num, &name)
	fmt.Println("Things read:", num, name)
	fmt.Println("Count of args scanned:", count, "Error:", err)
}

func readType3() {
	var num int
	var name string

	fmt.Print("Enter a number and string:")
	count, err := fmt.Scanln(&num, &name)
	fmt.Println("count:", count, "err:", err, "num:", num, "name:", name)
}

// For reading the whole line of text, use bufio
func readType4() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		if strings.Compare(text, "stop") == 0 {
			break
		}
		fmt.Println("Got:", text)
	}
}

func main() {
	//readType1()
	//readType2()
	//readType3()
	readType4()
}
