package main

import "fmt"

func g(f func(string)) {
	f("Hello from g!")
}

func main() {
	f := func(msg string) {
		fmt.Println(msg)
	}
	f("Hello!")
	g(f)
}
