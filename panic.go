package main

import "fmt"

func someFunction(n int) {
	if n > 3 {
		fmt.Println("OMG!")
		panic(n)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Relax...", r)
		}
	}()
	n := 0
	for {
		someFunction(n)
		n++
	}
	fmt.Println("This ain't gonna happen.")
}
