package printer

import (
	"fmt"
	"info/lib"
)

func init() {
	fmt.Println("printer.init(): Initializing printer package")
}

func Printit() {
	fmt.Println("Printing all languages: ")
	for _, lang := range lib.GetAll() {
		fmt.Println(lang)
	}
}
