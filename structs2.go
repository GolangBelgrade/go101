package main

import "fmt"

// START OMIT
type Address struct {
	Street string
}
type Person struct {
	Addr Address
	Name string
}

// END OMIT
func main() {
	pera := Person{
		Address: Address{"Savska 5"},
		Name:    "Pera",
	}
	fmt.Println(pera)
	fmt.Println(pera.Street)
}
