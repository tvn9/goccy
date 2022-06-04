package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	// set up a wait group
	var wg sync.WaitGroup
	// create a slice
	months := []string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	wg.Add(len(months))
	for i, x := range months {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()

	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
