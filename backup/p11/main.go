// Buffered Channels code example
package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a go data message
		i := <-ch
		fmt.Printf("Got %d from channel\n", i)

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		// the first 10 times through this loop, things go quickly;
		// after that, things slow down.
		fmt.Printf("sending %d to channel...\n", i)
		ch <- i
		fmt.Printf("sending %d to channel...\n", i)
	}

	fmt.Println("Done!")
	close(ch)
}
