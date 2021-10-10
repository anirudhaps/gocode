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
		// change below "\n" to "\r\n" if running on windows
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare(text, "stop") == 0 {
			break
		}
		fmt.Println("Got:", text)
	}
}

// reading a single UTF-8 unicode character
func readType5() {
	fmt.Print("Enter any utf-8 char:")
	reader := bufio.NewReader(os.Stdin)
	ch, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("char: %c\n", ch)
}

// using bufio's Scanner
func readType6() {
	fmt.Print("Enter input:> ")
	sc := bufio.NewScanner(os.Stdin)
	// the following will infinitely scan the input and read it and for every
	// valid input it will return true.
	// Returns: false in case of error or EOF
	// CTRL-D will generate EOF
	// Error can be seen by calling sc.Err(), for EOF, Err() returns <nil>
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	fmt.Println("Error stat:", sc.Err())
}

func main() {
	//readType1()
	//readType2()
	//readType3()
	//readType4()
	//readType5()
	readType6()
}
