package main

import "fmt"

// START1 OMIT
type Thing interface {
	Hello() string
}

func Greeting(t Thing) {
	fmt.Println(t.Hello())
}

// END1 OMIT
// START2 OMIT
type Person struct {
	Name string
}

func (p Person) Hello() string {
	return "Hello from " + p.Name
}

type Animal struct {
	Type, Name string
}

func (a *Animal) Hello() string {
	return "Hello from " + a.Name + " the " + a.Type
}

//END2 OMIT
func main() {
	p := Person{"Finn"}
	a := &Animal{"dog", "Jake"}
	Greeting(p)
	Greeting(a)
}
