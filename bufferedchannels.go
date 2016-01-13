package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("Sent", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Got", <-ch)
	}
}
