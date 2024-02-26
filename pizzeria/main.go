package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// type variables
const NumberOfPissas = 10

// number of pizzasMade, pizzasFieled, and total
var pizzasMade, pizzasFailed, total int

// Producer struct holds data of type chan PizzaOrder struct and quit chan error
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder struct holds pizzaNumber, message, and success bool
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPissas {
		fmt.Printf("Received order #%d!\n", pizzaNumber)
		delay := rand.Intn(5) + 1
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		// delay for sometime
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		// populate the PizzaOrder struct with above processed values
		pizza := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		// return the PizzaOrder object to the called function
		return &pizza
	}

	// if the pizzaNumber came in above the limit, no new pizza will be made
	// and the existing pizzaNumber in PizzaOrder object will be returned.
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza being made
	i := 0

	// this loop will continue to execute, trying to make pizzas,
	// until the quit channel receives something.
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// tried to make a pizza (sent something to data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run the consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPissas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 6:
		color.Red("It was an awful day...")
	case pizzasFailed >= 4:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 3:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 1:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was an awsome day!.")
	}
}
