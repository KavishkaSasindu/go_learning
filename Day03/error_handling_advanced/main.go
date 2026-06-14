package main

import (
	"errors"
	"fmt"
)

// control signal errors
var EOF error = errors.New("EOF")

func ReadFile() error {
	failed := false

	if failed {
		return errors.New("Failed to read")
	} else {
		return EOF
	}
}

// join errors , so if we have multiple errors belongs to one we can join the errors
func joinErrors() error {
	error_one := errors.New("Error one")
	error_two := errors.New("Error two")

	join_error := errors.Join(error_one, error_two)
	return join_error
}

func main() {
	fmt.Println("Hello World")

	// basically there is a built in package called errors and it has methods.Look below example
	new_error := errors.New("My new error")
	fmt.Println(new_error)

	// join errors example
	join_error := joinErrors()
	fmt.Println(join_error)

}
