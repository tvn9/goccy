package main

import "time"

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

}
