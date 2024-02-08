package main

// Example of functions that take arguments can run as goroutines

import (
	"fmt"
	"time"
)

func f(s string) {
	fmt.Printf("Goroutine %s\n", s)
}

func main() {
	for _, s := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		go f(s)
	}
	time.Sleep(1 * time.Second)
}
