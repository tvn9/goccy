package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously.

const eatTurn = 5

var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second
var delay = 1 * time.Second
var wg sync.WaitGroup
var myTurn sync.Mutex
var orderFinished []string
var orderMutex sync.Mutex

func diningProblem(philosopher string, myTurn *sync.Mutex) {
	defer wg.Done()

	fmt.Printf("%q is seated\n", philosopher)
	time.Sleep(delay)

	for i := eatTurn; i > 0; i-- {
		fmt.Printf("%q turn to eat.\n", philosopher)
		time.Sleep(delay)

		myTurn.Lock()
		fmt.Printf("%q start eating.\n", philosopher)
		time.Sleep(eatTime)
		myTurn.Unlock()

		fmt.Printf("%q is thinking.\n", philosopher)
		time.Sleep(thinkTime)
	}

	fmt.Printf("%q finishs eating satified.\n", philosopher)
	time.Sleep(delay)

	fmt.Printf("%q has left the table.\n", philosopher)
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()
}

func main() {
	// making list of philosopher
	philosophers := []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}

	wg.Add(len(philosophers))

	// philosopher taking turn to eat
	for i := 0; i < len(philosophers); i++ {
		// first philosopher start eating
		go diningProblem(philosophers[i], &myTurn)
	}
	wg.Wait()

	fmt.Println("The table is empty.")
	fmt.Println("-------------------")

	fmt.Printf("Order finished: %q\n", strings.Join(orderFinished, ", "))
}
