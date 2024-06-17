package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).
	// Taking the long way, it might have been written as:
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}
