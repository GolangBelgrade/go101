package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	a := Person{"Pera"}
	fmt.Println(a.Name)

	p := &a
	fmt.Println(p.Name)

	fmt.Println(p.Name == a.Name)
}
