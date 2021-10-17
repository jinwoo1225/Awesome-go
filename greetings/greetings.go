package greetings

//https://golang.org/doc/tutorial/create-module

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
