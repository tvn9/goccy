package main

import (
	"fmt"
	"sync"
)

// WaitGroup example

func greeting(s string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello %s - from %d\n", s, id)
}

func main() {
	const num = 5

	// define a wait group
	var wg sync.WaitGroup

	// define the total number of wait group
	wg.Add(num)
	for i := 0; i < num; i++ {
		go greeting("Thanh", i+1, &wg)
	}
	wg.Wait()
}
