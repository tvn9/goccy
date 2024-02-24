package main

import (
	"fmt"
	"sync"
)

// Example of goroutine with code that fix the problem of unsure join point
// in ex1

var wg sync.WaitGroup

func main() {
	yourName := func(n string) {
		defer wg.Done()
		fmt.Println(n)
	}

	wg.Add(2)
	go sayHello()
	go yourName("Thanh")
	wg.Wait()
}

func sayHello() {
	defer wg.Done()
	fmt.Print("hello, ")
}
