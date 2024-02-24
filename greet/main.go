// Start a goroutine with Go keyword.
package main

import (
	"fmt"
	"time"
)

func printGreeting(s string) {
	greet := "Hello,"
	fmt.Printf("%s %s!\n", greet, s)
}

// in Go main function is also a goroutine
func main() {
	name := "Thanh"

	// this is how a function spin off into it own goroutine
	// and no longer in sync with main goroutine, the function
	// might complete after main already finish, and might not
	// be able to show any output on the screen.
	go printGreeting(name)

	// this delay timing slow down main so that other goroutine
	// has a change to run and finish before main.
	// This however is not recommended in realwork application.
	time.Sleep(1 * time.Second)

	printGreeting("Mike")
}
