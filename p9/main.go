// Introduction to channels
package main

import (
	"fmt"
	"strings"
)

// Ping-Pong example code that will send sign from ping chan and receive from pong chan
func shout(ping <-chan string, pong chan<- string) {
	for {
		//
		s, ok := <-ping
		if !ok {
			// do something
		}
		pong <- fmt.Sprintf("%s!!!\n", strings.ToUpper(s))
	}
}

func main() {
	// create two channels.
	ping := make(chan string)
	pong := make(chan string)

	// start a goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press Enter (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			// jump out of the loop
			break
		}

		// send userInput to "ping" channel
		ping <- userInput

		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")
	// close the channels
	close(ping)
	close(pong)
}
