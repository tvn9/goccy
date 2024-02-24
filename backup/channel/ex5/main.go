package main

// Buffered channel example

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Printf("Got %d from channel\n", i)

		// simulate doing a lot of work
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 30)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Printf("sending %d to channel...\n", i)
		ch <- i
		fmt.Printf("sent %d to channel!\n", i)
	}

	fmt.Println("Done!")
	close(ch)
}
