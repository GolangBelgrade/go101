package main

import "fmt"

// START OMIT
type Person struct {
	Name, Gender string
	Age          int
}

func (p *Person) Hello() {
	fmt.Println(fmt.Sprintf("%s says hi!", p.Name))
}

// END OMIT
func main() {
	pera := &Person{"Pera", "male", 25}
	pera.Hello()
}
