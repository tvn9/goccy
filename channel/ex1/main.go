// Channel example
package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			fmt.Println("ping channel problem.")
			return
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// Print a CLI prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		lowerQ := strings.ToLower(userInput)

		if lowerQ == "q" {
			break
		}

		ping <- userInput
		// wait for a respone
		respone := <-pong
		fmt.Println("Response:", respone)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
