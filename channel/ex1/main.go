package main

import "fmt"

func main() {
	chn := make(chan bool)
	go func() {
		chn <- true
	}()
	go func() {
		var y bool
		y <- chn
		fmt.Println(y)
	}()
}
