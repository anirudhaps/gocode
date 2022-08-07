package main

import "fmt"

func typeAssertionDemo(t interface{}) {
	// This is called type assertion where we try to assert what is the
	// underlying type of an interface. Notation: variable_name.(T) where
	// T is the underlying type. Type assertions can only take place on
	// interfaces. Type assertions are performed at runtime.

	// We may not know the underlying type of an interface with
	// surity. So, the second return value (ok below) tells us whether the type
	// assertion is correct or not. ok == true if correct, false otherwise.
	// When ok == false, the value is zero value of respective type.
	val, ok := t.(string)
	if ok {
		fmt.Printf("Underlying type is string, value is %s\n", val)
	} else {
		fmt.Printf("Underlying type is NOT string, value is empty: length: %v\n", len(val))
	}

	intVal, ok := t.(int)
	if ok {
		fmt.Printf("Underlying type is int, value is %d\n", intVal)
	} else {
		fmt.Printf("Underlying type is NOT int, value is zero: %d\n", intVal)
	}

	// we can use the type switch to check for the underlying type of an
	// interface if we are unsure about it.
	switch t.(type) {
	case string:
		fmt.Printf("Type switch: this is string type, value is: %s\n", t)
	case int:
		fmt.Printf("Type switch: this is int type, value is: %d\n", t)
	default:
		fmt.Printf("Type switch: type unknown\n")
	}
}

func main() {
	var a int = 5
	typeAssertionDemo(a)
	var b string = "hello"
	typeAssertionDemo(b)
}
