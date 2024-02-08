package main

// Example that demonstrates how multiple channels can be used to coordinate work.

import (
	"fmt"
	"math/rand"
)

type Work struct {
	value int
}

type Result struct {
	value int
}

func main() {
	// define 3 channels
	workCh := make(chan Work)
	resultCh := make(chan Result)
	done := make(chan bool)

	// Create a slice of Work struct
	workQueue := make([]Work, 100)
	for i := range workQueue {
		// assign and randome integer to 100 workQueue slice
		workQueue[i].value = rand.Int()
	}
	// Create 10 worker goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// Get work from the work channel
				work := <-workCh
				// Compute result
				result := Result{
					// take work.value multiply by 2 then assign
					// the result to Result struct value field
					value: work.value * 2,
				}
				// Send the result struct via the result channel
				resultCh <- result
			}
		}()
	}
	// Create a results slice of Result struct
	results := make([]Result, 0)
	go func() {
		// Collect all the results.
		for i := 0; i < len(workQueue); i++ {
			// take Result struct data from resultCh and append to
			// results slice of Result.
			results = append(results, <-resultCh)
		}
		// When all the results are collected, notify the done channel
		done <- true
	}()

	// Send all the work to the workers
	for _, work := range workQueue {
		workCh <- work
	}
	// Wait until everthing is done
	<-done

	// iterate the results and print all values
	for i, v := range results {
		fmt.Println(i, v.value)
	}
}
