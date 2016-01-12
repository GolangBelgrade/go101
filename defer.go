package main

import "fmt"

func main() {
	defer fmt.Println("Defer statement")
	fmt.Println("Regular statement")
}
