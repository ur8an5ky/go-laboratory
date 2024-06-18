package main

import (
	"fmt"
	"strings"

	"calculator.com/calculator/operations"
)

func main() {
	var val1, val2 int
	// var result float32
	var operation string

	fmt.Print("Type in the first number: ")
	fmt.Scan(&val1)

	fmt.Print("Type in the second number: ")
	fmt.Scan(&val2)

	fmt.Print("Type in the operation name or its number (1.add, 2.subtract, 3.multiply, 4.divide): ")
	fmt.Scan(&operation)

	operation = strings.ToLower(operation)

	fmt.Printf("The result of %s operation of given numbers (%d, %d) is ", operation, val1, val2)

	switch operation {
	case "1", "1.", "add":
		fmt.Print(val1 + val2)
	case "2", "2.", "subtract":
		fmt.Print(val1 - val2)
	case "3", "3.", "multiply":
		fmt.Print(val1 * val2)
	case "4", "4.", "divide":
		fmt.Print(float32(val1) / float32(val2))
	// add cases for factorial (how to treat the second value then), arithmetic mean etc.
	// also add error handling
	default:
		fmt.Print(operations.Factorial(val1))
	}

	fmt.Printf("\n")
}
