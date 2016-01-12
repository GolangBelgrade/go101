package main

import "fmt"
import "errors"

func someFunction() (int, string, error) {
	return 7, "hello", errors.New("ohnoez")
}

func main() {
	a, b, err := someFunction()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(err)

	_, c, _ := someFunction()
	fmt.Println(c)
}
