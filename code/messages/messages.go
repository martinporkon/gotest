package messages

// TODO the foal is to test theese functions

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!\n", name)
}

func depart(name string) string { // this is not exported as it starts with a non capital letter
	return fmt.Sprintf("Goodbye, %v\n", name)
}
