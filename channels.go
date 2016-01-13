package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		message <- "Hello!"
	}()
	fmt.Println("Waiting...")

	result := <-message
	fmt.Println(result)
}
