// Goroutine with waitgroup
package main

import (
	"fmt"
	"sync"
)

func printing(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s\n", s)
}

func main() {
	var wg sync.WaitGroup

	months := []string{"January", "Fubuary", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	wg.Add(len(months))
	for i, m := range months {
		go printing(fmt.Sprintf("#%-5d %s", i+1, m), &wg)
	}
	wg.Wait()

	wg.Add(1)
	printing("Some thing else to say!", &wg)
	wg.Wait()
}
