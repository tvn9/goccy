package main

import (
	"fmt"
	"sync"
)

// This example will deterministically block the main goroutine until the
// goroutine hosting the sayHello function terminated.
// Learning the sync.WaitGroup in the sync package

var wg sync.WaitGroup

func main() {
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}
