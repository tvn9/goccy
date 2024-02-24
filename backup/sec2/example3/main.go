package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	words := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"November",
		"December",
	}
	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("#%d, %v", i, x), &wg)
	}
	wg.Wait()
	printSomething("We already in Dec, 2023.", &wg)
}
