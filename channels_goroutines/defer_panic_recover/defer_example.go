package main

import "fmt"

// defer keyword is used to defer (postpone) the execution of a function until the
// surrounding/calling function returns. This is mostly used to perform some cleanup just after
// surrounding/calling function returns.
//
// The defer statement (a deferred func call) pushes the function call into a list of saved calls
// (a stack). Just after returning from the surrounding function, the stack is popped and the call
// is invoked.

func printer(value string) {
	fmt.Println(value)
}

func main() {
	// due to defer keyword, this printer call will be not be executed now. It will be executed
	// just after the main function returns. See by running this program.
	defer printer("Returning from main") // statement-1
	initial := 0

	// There can be many defered function calls. Two more are invoked below.
	// Important property of defer statements:
	// A deferred function's arguments are evaluated when the defer statement is evaluated.
	// It means this defer statement is evaluated now. initial is evaluated now and 0 is placed in
	// the call as argument. Thus, when this deferred call is actually invoked (when main returns),
	// it prints "Initial value of counter was 0".
	defer fmt.Println("Initial value of counter was", initial) // statement-2

	initial = 3
	// this call is evaluated now and 3 is passed as argument to this call. Even though, initial
	// becomes 5, this call will be invoked with initial=3.
	defer func(j int) {
		for j < 5 {
			fmt.Println("defered:", j)
			j++
		}
	}(initial) // statement-3

	initial = 5

	// Because each of the defer function calls are pushed into stack, these are executed in
	// LIFO order (when the surrounding function, main() here, returns). For example, statement-1
	// is pushed into the stack of deferred calls, then statement-2, then statement-3 is pushed.
	// When main returns, statement-3 is first executed, then statement-2, then statement-1 (LIFO).

	fmt.Println("Demonstrate LIFO order of deferred calls execution: ")
	lifoOrder()
	fmt.Println()

	if err := validateViewerEligibility("A", 12); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Main here....")
}

func lifoOrder() {
	// another example to demonstrate LIFO order execution of deferred calls
	for i := 0; i < 4; i++ {
		defer fmt.Println("lifo:", i) // prints 3210
	}
}

// Another property of defer statement:
// Deferred functions may read and assign to the returning function's named return values.
// err is the named error
func validateViewerEligibility(movieType string, viewerAge uint32) (err error) {
	defer func() {
		// this deferred function updates the returned error message after the surrounding function
		// (validateViewerEligibility) returns
		if err != nil {
			err = fmt.Errorf("Viewers of age %v are ineligible to view movies of type %v", viewerAge, movieType)
		}
	}()

	if movieType == "A" && viewerAge < 18 {
		err = fmt.Errorf("Error")
		return err
	}
	return nil
}
