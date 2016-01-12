package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["x"] = "foo"
	m["y"] = "bar"
	fmt.Println(m)

	n := map[string]string{
		"x": "foo",
		"y": "bar",
	}
	fmt.Println(n)
}
