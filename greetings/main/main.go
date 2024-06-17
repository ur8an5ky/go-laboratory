package main

import (
	"fmt"
	"greets/greetings"
)

func main() {
	message := greetings.Hello("Kuba")
	fmt.Println(message)
}
