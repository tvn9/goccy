package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
	printMessage()
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {
	// challenge: modify this code so that the calls to updateMessage() on lines
	// 27, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	var str []string

	str1 := "Hello, universe!"
	str2 := "Hello, cosmos!"
	str3 := "Hello, world!"

	str = append(str, str1, str2, str3)

	wg.Add(len(str))
	for i, s := range str {
		updateMessage(fmt.Sprintf("Test %d %s", i, s))
	}
	wg.Wait()
}
