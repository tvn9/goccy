package main

import (
	"fmt"
	"time"
)

// Race Conditions - Data race

func main() {
	var data int

	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
