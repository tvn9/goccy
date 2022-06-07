package main

import (
	"fmt"
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

const eatTurns = 3

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var delay = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second

func diningProblem(philosophers string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	fmt.Printf("%q is seated.\n", philosophers)
	time.Sleep(delay)

	for i := eatTurns; i > 0; i-- {
		fmt.Printf("%q turn to eat.\n", philosophers)
		time.Sleep(delay)

		leftFork.Lock()
		fmt.Printf("\t%s pickup right fork.\n", philosophers)
		rightFork.Lock()
		fmt.Printf("\t%s pickup left fork.\n", philosophers)

		fmt.Printf("%q has both forks to eat.\n", philosophers)
		time.Sleep(eatTime)

		fmt.Printf("%q is thinking.\n", philosophers)
		time.Sleep(thinkTime)

		rightFork.Unlock()
		leftFork.Unlock()
		fmt.Printf("\t%q putdown leftFork and rightFork.\n", philosophers)
		time.Sleep(delay)
	}

	//
	fmt.Printf("%q is full and satisfied.\n", philosophers)
	time.Sleep(delay)

	fmt.Printf("%q has left the table.\n", philosophers)
}

func main() {
	fmt.Println("The Dinning Philosophers Problem")
	fmt.Println("--------------------------------")

	wg.Add(len(philosophers))

	forkLeft := &sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {
		forkRight := &sync.Mutex{}

		go diningProblem(philosophers[i], forkLeft, forkRight)

		forkLeft = forkRight
	}
	wg.Wait()

	fmt.Println("The table is empty.")
}
