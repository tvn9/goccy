package main

// Basic example of how to create a goroutine

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go f()

	fmt.Println("Hello from main!")
	time.Sleep(1 * time.Second)
}
