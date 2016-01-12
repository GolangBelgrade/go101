package main

import "fmt"

func main() {
	s := make([]string, 0)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, "foo")
	s = append(s, "bar")
	fmt.Println(len(s))
	fmt.Println(cap(s))

	c := make([]string, 0)
	copy(c, s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
