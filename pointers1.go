package main

import "fmt"

func main() {
	a := "foo"

	p := &a // pointer to a
	fmt.Println(p)

	v := *p // value from a pointer
	fmt.Println(v)

	fmt.Println(v == a)
}
