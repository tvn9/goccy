package main

import (
	"fmt"
	"runtime"
	"sync"
)

// In this example, we combine the fact that goroutines are not gargage
// collected with the runtime's ability to introspect upon itself and measure
// the amount of memory allocated before and after goroutine creation:

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	// 1. Require a goroutine that will never exit
	noop := func() {
		wg.Done()
		<-c
	}
	// 2. define the number of goroutines to create
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	// 3. measure the amount of memory consumed before creating goroutines
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	// 4. measure the amount of memory consumed after creating goroutines
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
