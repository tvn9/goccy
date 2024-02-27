package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Philosopher struct holds name leftFork, and rightFork
type Philosopher struct {
	name                string
	leftFork, rightFork int
}

// philosophers is the list of all philosophers
var philosophers = []Philosopher{
	{name: "Washington", leftFork: 4, rightFork: 0},
	{name: "Adams", leftFork: 0, rightFork: 1},
	{name: "Jefferson", leftFork: 1, rightFork: 2},
	{name: "Madison", leftFork: 2, rightFork: 3},
	{name: "Monroe", leftFork: 3, rightFork: 4},
}

// variables for delay time
var (
	eatTurn   = 3 // how many times each philosopher eatTurn
	thinkTime = time.Millisecond * time.Duration(rand.Intn(1000))
	eatTime   = time.Second * time.Duration(rand.Intn(3))
	sleepTime = time.Millisecond * time.Duration(rand.Intn(1000))
)

// *** challenge solution ***
var orderMutex sync.Mutex
var orderFinished []string

// main is where the program start
func main() {
	// Print out a welcome message
	fmt.Println("Dinning Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty.")

	// *** challenge solution ***
	time.Sleep(sleepTime)

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")

	// Order finished

	fmt.Println("Order finished:")
	for i, phi := range orderFinished {
		fmt.Printf("#%d: %s\n", i+1, phi)
	}
}

func dine() {
	/*
		eatTime = 0 * time.Second
		thinkTime = 0 * time.Second
		sleepTime = 0 * time.Second
	*/

	wg := sync.WaitGroup{}
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
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], &wg, forks, seated)
	}
	wg.Wait()
}

func diningProblem(phi Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated a the table.\n", phi.name)
	seated.Done()

	seated.Wait()

	// eat three time
	for i := eatTurn; i > 0; i-- {
		if phi.leftFork > phi.rightFork {
			forks[phi.rightFork].Lock()
			fmt.Printf("\t%s pickup the rightFork.\n", phi.name)
			forks[phi.leftFork].Lock()
			fmt.Printf("\t%s pickup the leftFork.\n", phi.name)
		} else {
			forks[phi.leftFork].Lock()
			fmt.Printf("\t%s pickup the leftFork.\n", phi.name)
			forks[phi.rightFork].Lock()
			fmt.Printf("\t%s pickup the rightFork.\n", phi.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", phi.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", phi.name)
		time.Sleep(thinkTime)

		forks[phi.leftFork].Unlock()
		forks[phi.rightFork].Unlock()
		fmt.Printf("\t%s put down the forks.\n", phi.name)

	}
	// *** challenge solution ***
	orderMutex.Lock()
	orderFinished = append(orderFinished, phi.name)
	orderMutex.Unlock()

	fmt.Printf("%s is satisfying.\n", phi.name)
	fmt.Printf("%s left the table.\n", phi.name)
}
