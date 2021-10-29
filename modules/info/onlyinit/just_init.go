package onlyinit

import "fmt"

func init() {
	fmt.Println("onlyinit.init(): invoked")
}

func WontCall() {
	fmt.Println("WontCall() invoked")
}
