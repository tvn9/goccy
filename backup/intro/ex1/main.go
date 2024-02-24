package main

import (
	"fmt"
	"time"
)

// Example of how to start go routine

func main() {
	go sayHello()
	// Doing other things

	// start go concurrency with anonymous function
	go func() {
		fmt.Println("Goodbye!")
	}()

	// or using variable func
	sayYourName := func(name string) {
		fmt.Printf("Hello, %s\n", name)
	}

	go sayYourName("Thanh")
	time.Sleep(1 * time.Second)
}

func sayHello() {
	fmt.Println("Hello!")
}
