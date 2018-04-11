package main

import (
	"fmt"
	"os"
)

func main() {
	clargs := os.Args
	if len(clargs) != 2 {
		fmt.Println("Error: Too many command line arguments")
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
	file, err := os.Open(clargs[1])
	if err != nil {
		fmt.Println("Error during file opening")
		os.Exit(1)
	}
	data := make([]byte, 99999)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println("Error during file reading")
		file.Close()
		os.Exit(1)
	}
	fmt.Println(string(data[0:n]))
}
