package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Jinwoo")
	fmt.Println(message)
}
