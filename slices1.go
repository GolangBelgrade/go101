package main

import "fmt"

func main() {
	var s []string              // slice of string declaration
	s = make([]string, 3)       // empty allocation
	s = []string{"1", "2", "3"} // explicit allocation
	fmt.Println(len(s))
	fmt.Println(s)

	s[0] = "x"
	s[1] = "y"
	s[2] = "z"
	fmt.Println(len(s))
	fmt.Println(s)
}
