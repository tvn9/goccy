package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Channels are first-class values, just like strings or integers.

func boring(msg string) <-chan string { // Return recive-only channel of strings.
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func main() {
	c := boring("boring!")
	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Printf("You are boring; I'm leaving.")
}
