package main

import "fmt"

// In this example, the write to x is sequenced before the channel write, which is synchronied
// before the channel read. The printing is sequenced after the channel read, so the write to x
// happens before the printing of X.
func main() {
	var x int
	ch := make(chan int)
	go func() {
		<-ch
		fmt.Println(x)
	}()

	x = 1
	ch <- 0
	select {}
}
