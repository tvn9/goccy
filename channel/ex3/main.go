// channel with buffer example
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

	}
}

func main() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel...")
		ch <- i
	}
}
