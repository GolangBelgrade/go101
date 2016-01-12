package main

import (
	"fmt"
	"log"
)

type City struct {
	name string
}

func (c *City) Hello() string {
	return fmt.Sprintf("Hello %s", c.name)
}

func main() {
	city := &City{"Belgrade"}
	greeting := city.Hello()

	log.Println(greeting)
}
