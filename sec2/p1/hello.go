package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {

	// This is only an intro to go example, not a productio way of
	// coding real go concurrency.
	go printSomething("Print something to the screen!")

	time.Sleep(1 * time.Second) // Bad way of implement goroutines

	printSomething("Print this second thing on the screen!")
}
