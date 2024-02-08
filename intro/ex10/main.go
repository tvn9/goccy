package main

import (
	"fmt"
	"time"
)

func main() {
	for _, s := range []string{"A", "B", "C", "D", "E", "F", "G"} {
		s := s
		go func(s string) {
			fmt.Printf("Goroutine %s\n", s)
		}(s)
		time.Sleep(1000)
	}
}
