package main

import (
	"fmt"
	"time"
)

func main() {
	// There is a data race problem in this for loop block ther is a shared variable that is
	// written by on goroutine and read by 7 others without any synchronization.
	for _, s := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		go func() {
			fmt.Printf("Group 1: Goroutine %s\n", s)
		}()
	}

	go func() {
		for _, a := range []string{"A", "B", "C", "D", "E", "F", "G"} {
			fmt.Printf("Group 2: Goroutine %s\n", a)
		}
	}()
	time.Sleep(1 * time.Second)
}
