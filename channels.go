package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Hello!"
	}()
	fmt.Println("Waiting...")

	result := <-ch
	fmt.Println(result)
}
