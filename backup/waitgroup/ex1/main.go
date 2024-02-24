package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup is a great way to wait for a set of concurrent operations to complete
// when you either don't care about the result of the concurrent operation, or you have
// other means of collecting their results.

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()
	wg.Wait()
	fmt.Println("All goroutines completed.")
}
