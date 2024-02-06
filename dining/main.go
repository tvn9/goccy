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
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal.
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosophers
		go diningProblem(philosophers[i], wg, forks, seated)
	}
}

func diningProblem(phi Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seat *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", phi.name)
	seat.Done()

	seat.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks
		forks[phi.leftFork].Lock()
		fmt.Printf("\t%s takes the left fork.\n", phi.name)
		forks[phi.rightFork].Lock()
		fmt.Printf("\t%s takes the right fork.\n", phi.name)

		fmt.Printf("\t%s hask both forks and is eating.\n", phi.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s philosopher is thinking.\n", phi.name)
		time.Sleep(thinkTime)

		forks[phi.leftFork].Unlock()
		forks[phi.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", phi.name)
	}
	fmt.Println(phi.name, "is satisfied.")
	fmt.Println(phi.name, "left the table.")
}
