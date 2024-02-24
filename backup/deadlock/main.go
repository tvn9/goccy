package main

// Deadlock example
// A deadlock program is the result of all concurrent process are waiting on one another.
// in this state, the program will never recover without outside intervention.

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func main() {
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock() // v1 is lock
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock() // v2 is also lock
		defer v2.mu.Unlock()
		// here is the  demonstration of a deadlock condition

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b) // here both goroutine are waiting infinitely on each other.
	go printSum(&b, &a)
	wg.Wait()
}
