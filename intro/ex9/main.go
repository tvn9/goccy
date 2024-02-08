package main

// Another exemple of to demonstrate and data race condition

import (
	"fmt"
	"time"
)

func main() {

	var s string
	s = "A"
	go func() {
		fmt.Printf("Goroutine %s\n", s)
	}()

	s = "B"
	go func() {
		fmt.Printf("Goroutine %s\n", s)
	}()

	s = "C"
	go func() {
		fmt.Printf("Goroutine %s\n", s)
	}()

	time.Sleep(1000)
}
