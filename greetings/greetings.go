package greetings

//https://golang.org/doc/tutorial/create-module

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greetings for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init settings initial values
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greetings messages.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
		"안녕? %v, 만나서 반갑다!",
	}

	// Returns a randomly selected message format by specifying
	// a random index for the slice of formats
	return formats[rand.Intn(len(formats))]
}
