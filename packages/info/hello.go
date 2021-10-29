package main

import (
	"fmt"
	_ "info/onlyinit"
)

// Note: the _ above (in import statement) is called a blank identifier.
// We give a blank identifier to an (to be) imported package if we only
// want to invoke its init() function and no subsequent usage in the code
// If we don't mark such a package with _, the golang build system will
// throw "package imported but not used" compile time error

// the above way of importing package is also sometimes called as:
// Importing a package only for its side effects

func hello() {
	fmt.Println("hello")
}

func init() {
	fmt.Println("hello.go: init(): invoked")
}
