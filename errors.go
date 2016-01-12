package main

import (
	"errors"
	"log"
)

func someFunction() (string, error) {
	return "foo", errors.New("bar")
}

func main() {
	_, err := someFunction()
	if err != nil {
		log.Printf("got error: %s", err)
	}
}
