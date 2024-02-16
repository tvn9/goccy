package main

import (
	"fmt"
	"strings"
)

// Base channel example

func shout(send <-chan string, receive chan<- string) {
	for {
		s := <-send
		receive <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	send := make(chan string)
	receive := make(chan string)

	// start a goroutine and call shout(chan string, chan string)
	go shout(send, receive)

	fmt.Println("Type something then press ENTER (enter Q to quit)")
	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		fmt.Scanln(&userInput)

		if "q" == strings.ToLower(userInput) {
			break
		}

		send <- userInput

		// wait for a response
		response := <-receive

		fmt.Println("Response:", response)
	}
	fmt.Println("All done. Closing channels.")
	close(send)
	close(receive)
}
