package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
	printMessage()
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {
	msgs := []string{}

	str0 := "Hello, universe!"
	str1 := "Hello, cosmos!"
	str2 := "Hello, world!"

	msgs = append(msgs, str1, str2, str0)

	wg.Add(len(msgs))
	for _, m := range msgs {
		go updateMessage(m, &wg)
	}
	wg.Wait()

}

/*
func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!", &wg)
	wg.Wait()
	printMessage()
}
*/
