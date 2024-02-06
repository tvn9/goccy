package main

import (
	"fmt"
	"sync"
	"time"
)

// The dining philosophers problems

// philosopher struct stores information about the philosopher
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers is a list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define some variables
var (
	hunger    = 3 // how many times does a person eat?
	eatTime   = 1 * time.Second
	thinkTime = 3 * time.Second
	sleepTime = 1 * time.Second
)

func main() {
	// print out a welcome message
	fmt.Println("Dinning Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")
}

func dine() {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	// wait group
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// Wait for everyone to be seated before eating
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks, forks are fields leftFork and rightFork in
	// Philosophers	struct
	var forks = make(map[int]*sync.Mutex) // map of forks's key 0-4, and type sync.Mutex
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{} // assign key 0-4 & type sync.Mutex to forks map
	}

	// start the meal by iterating through the slice of Philosophers
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosophers
		go diningRoutine(philosophers[i], wg, forks, seated)
	}

	// Wait for the Philosophers to finish. This blocks until the waigGroup is 0.
	wg.Wait()
}

func diningRoutine(phi Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seat *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", phi.name)

	// Decrement the seated WaitGroup by one.
	seat.Done()

	// Wait until everyone is seated
	seat.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		if phi.leftFork > phi.rightFork {
			// get a lock on both forks
			forks[phi.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", phi.name)
			forks[phi.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", phi.name)
		} else {
			forks[phi.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", phi.name)
			forks[phi.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", phi.name)
		}

		// the philosopher has both forks locked, and start eating
		fmt.Printf("\t%s hask both forks and is eating.\n", phi.name)
		time.Sleep(eatTime)

		// the philosopher starts to think, and still has the forks
		fmt.Printf("\t%s philosopher is thinking.\n", phi.name)
		time.Sleep(thinkTime)

		// unlock the mutexes for both forks
		forks[phi.leftFork].Unlock()
		forks[phi.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", phi.name)
	}
	// the philosopher has finished eating.
	fmt.Println(phi.name, "is satisfied.")
	fmt.Println(phi.name, "left the table.")
}
