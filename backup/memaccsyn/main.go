package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var data int
	// data = 1
	go func() {
		memoryAccess.Lock() // Not recommended for real project
		data++
		memoryAccess.Unlock() // only a demonstration memory access synchronization
	}()
	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("The value is %v\n", data)
	} else {
		fmt.Printf("The value is %v\n", data)
	}
	memoryAccess.Unlock()

	// While we have solved our data race, we haven't actually solved
	// the race condition.

	// In this example, either the goroutine will execute first, or both
	// the if and else block will. we don't know which will occur first
	// in any given execution

	// The order of this program is still nondeterministic.
}
