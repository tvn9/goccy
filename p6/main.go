package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	order chan PizzaOrder
	quit  chan chan error
}

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

func makePizza(pizNum int) *PizzaOrder {
	pizNum++
	if pizNum <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizNum)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizNum, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizNum)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizNum)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizNum)
		}

		p := PizzaOrder{
			pizzaNumber: pizNum,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizNum,
	}
}

func pizzeria(pm *Producer) {
	// keep track of pizza made
	var i = 0

	// run until user quit
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pm.order <- *currentPizza:
			case quitChan := <-pm.quit:
				// close channels
				close(pm.order)
				close(quitChan)
				return
			}
		}
	}
	// try to make pizza
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// printout a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		order: make(chan PizzaOrder),
		quit:  make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.order {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizza...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("fails to close channel.", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("-----------------")
	color.Cyan("done for the day.")
	color.Cyan("We make %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was a okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was a great day...")
	}
}
