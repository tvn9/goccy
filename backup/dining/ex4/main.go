package main

// This example demonstrate implementation with channel

import (
	"fmt"
	"math/rand"
	"time"
)

func philosopher(phi string, leftFork, rightFork chan bool) {
	// Each Philosopher get to eat 3 times.
	for i := 3; i > 0; i-- {
		// Think for some time
		fmt.Printf("Philosopher %s is thinking.\n", phi)
		time.Sleep(time.Duration(rand.Intn(1000)))
		select {
		case <-leftFork:
			select {
			case <-rightFork:
				fmt.Printf("Philosopher %s is eating\n", phi)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
				rightFork <- true
			default:
			}
			leftFork <- true
		}
	}
	fmt.Printf("Philosopher %s left the table.\n", phi)
}

func main() {
	phi := []string{"Socrates", "Aristotle", "Kant", "Descartes", "Nietzsche"}

	var forks [5]chan bool
	for i := range forks {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}

	go philosopher(phi[0], forks[4], forks[0])
	go philosopher(phi[1], forks[0], forks[1])
	go philosopher(phi[2], forks[1], forks[2])
	go philosopher(phi[3], forks[2], forks[3])
	go philosopher(phi[4], forks[3], forks[4])
	select {}
}
