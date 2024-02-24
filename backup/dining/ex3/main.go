package main

// This version example demonstrate implementation prone to deadlock condition

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func philosopher(phi string, firstFork, secondFork *sync.Mutex) {
	// Each Philosopher get to eat 3 times.
	for i := 3; i > 0; i-- {
		// Think for some time
		fmt.Printf("Philosopher %s is thinking.\n", phi)
		time.Sleep(time.Millisecond + time.Duration(rand.Intn(1000)))
		// Get left fork
		firstFork.Lock()
		fmt.Printf("Philosopher %s got the left fork\n", phi)
		// Get right fork
		secondFork.Lock()
		fmt.Printf("Philosopher %s get the right fork\n", phi)
		// Start eating
		fmt.Printf("Philosopher %s is eating\n", phi)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		secondFork.Unlock()
		firstFork.Unlock()
		fmt.Printf("Philosopher %s put down the forks\n", phi)
	}
	fmt.Printf("Philosopher %s left the table.\n", phi)
}

func main() {
	phi := []string{"Socrates", "Aristotle", "Kant", "Descartes", "Nietzsche"}
	forks := [5]sync.Mutex{}

	go philosopher(phi[0], &forks[0], &forks[4])
	go philosopher(phi[1], &forks[0], &forks[1])
	go philosopher(phi[2], &forks[1], &forks[2])
	go philosopher(phi[3], &forks[2], &forks[3])
	go philosopher(phi[4], &forks[3], &forks[4])
	select {}
}
