package main

// Note: the numbers (1: 2: etc.) in the below comments show order of execution
import (
	"fmt"
	"info/lib"     // 1: lib package is imported ==> lib.init() is invoked
	"info/printer" // 2: printer package is imported ==> printer.init() is invoked
)

var num = dummy() // 3: variable declaration's initializer (i.e. dummy()) is evaluated

func main() { // 7: main() is invoked
	lib.Add("hk", "Haskell")
	lib.Add("njs", "NodeJs")
	fmt.Println("New language: ", lib.Get("hk"))
	// the below statement will throw compile-time error:
	// "cannot refer to unexported name printer.init"
	// due to the fact that init is starting with small letter, so not exported
	// printer.init()
	printer.Printit()
	fmt.Println("num:", num)

	// a func, defined in some other file that lies in the same package,
	// can be invoked w/o package prefix
	hello()
}

// init() functions are defined in package blocks and are used for:
// - variables initialization if cannot be done in initialization expression
// - checking / fixing program’s state
// - registering
// - running one-time computations
// and many more

// 4: hello.go: init() is invoked because hello.go filename comes before main.go alphabetically
func init() { // 5: this is invoked i.e. the first init() function that appears in order
	fmt.Println("main.init() - 1: initializing")
}

func dummy() int32 { // 3: this is invoked
	fmt.Println("dummy(): invoked")
	return 43
}

func init() { // 6: this is invoked i.e. the second init() func that appears in order and so on
	fmt.Println("main.init() - 2: initializing")
}

// Note: There can be multiple init() functions in the file
// they will be executed in the order they show up in the file
// If they span multiple files, they will be executed in lexical filename order

// Important: Package initialization is done only once even if package is imported many times
// Package initialization includes all the above (numbered) steps
// Package initialization happens in the alphabetical order of the file-names
// For eg. both hello.go and main.go are in the same main package. All init() functions of hello.go
// will be invoked first and then all init() functions of main.go will be invoked.
// Tip: run this module and see the sequence of prints
