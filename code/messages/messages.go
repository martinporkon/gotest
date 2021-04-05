package messages

// TODO the foal is to test theese functions

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!\n", name)
}

func depart(name string) string {
	return fmt.Sprintf("Goodbye, %v\n", name)
}
