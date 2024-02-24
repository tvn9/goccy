package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(index int, done <-chan struct{}, output chan<- int) {
	for {
		// Produce a random value
		value := rand.Int()
		// Wait a bit, using random value for sleep time
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		// Send the value
		select {
		case output <- value:
		case <-done:
			return
		}
		fmt.Printf("Producer %d sent %d\n", index, value)
	}
}

func consumer(index int, input <-chan int) {
	for value := range input {
		fmt.Printf("Consumer %d received %d\n", index, value)
	}
}

func main() {
	doneCh := make(chan struct{})
	dataCh := make(chan int, 0)
	for i := 0; i < 10; i++ {
		go producer(i, doneCh, dataCh)
	}
	for i := 0; i < 10; i++ {
		go consumer(i, dataCh)
	}
	time.Sleep(time.Second * 10)
	close(doneCh)

}
