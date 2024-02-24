// Go Routines
package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("Hello Go Concurrency!")

	time.Sleep(time.Second * 1) // don't do this, this is the bad example of how to use go routine.
	printSomething("Hello Gopher!")

}
