package main

import "fmt"

// This program also print 1

func main() {
	var x int
	ch := make(chan int)
	go func() {
		ch <- 0
		fmt.Println(x)
	}()
	x = 1
	<-ch
	select {}
}
