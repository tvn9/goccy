package main

import (
	"fmt"
	"log"
	"strings"
)

// Base send and receive channels example

func shout(send <-chan string, receive chan<- string) {
	for {
		s, ok := <-send
		if !ok {
			log.Fatal("something wrong with send channel")
		}
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
		response, ok := <-receive
		if !ok {
			log.Fatal("something wrong with receive channel.")
		}

		fmt.Println("Response:", response)
	}
	fmt.Println("All done. Closing channels.")
	close(send)
	close(receive)
}
