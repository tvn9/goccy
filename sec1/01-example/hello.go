package main

import "fmt"

func printSomething(s string) {
	fmt.Println(s)
}

func main() {

	// This is only an intro to go example, not a productio way of
	// coding real go concurrency.
	go printSomething("Print something to the screen!")

	printSomething("Print this second thing on the screen!")
}
