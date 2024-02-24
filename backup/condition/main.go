package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Queue struct {
	elements    []int
	front, rear int
	len         int
}

// NewQuene initializes an empty circular queue with the given capacity
func NewQuene(capacity int) *Queue {
	return &Queue{
		elements: make([]int, capacity),
		front:    0,
		rear:     -1,
		len:      0,
	}
}

// Enqueue adds a value to the queue. Returns false if queue is full
func (q *Queue) Enqueue(value int) bool {
	if q.len == len(q.elements) {
		return false
	}
	// Advance the write pointer, go around in circle
	q.rear = (q.rear + 1) % len(q.elements)
	// Write the value
	q.elements[q.rear] = value
	q.len++
	return true
}

// Dequeue removes a value from the queue. Returns 0, false if queue is empty
func (q *Queue) Dequeue() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	// Read the value at the read ponter
	data := q.elements[q.front]
	// Advance the read pointer, go around in circle
	q.front = q.elements[q.front]
	// Advance the read pointer, go around in a circle
	q.front = (q.front + 1) % len(q.elements)
	q.len--
	return data, true
}

// The starting point
func main() {
	lock := sync.Mutex{}
	fullCond := sync.NewCond(&lock)
	emptyCond := sync.NewCond(&lock)
	queue := NewQuene(10)
	producer := func() {
		for {
			// Produce value
			value := rand.Int()
			lock.Lock()
			for !queue.Enqueue(value) {
				fmt.Println("Queue is full")
				fullCond.Wait()
			}
			lock.Unlock()
			emptyCond.Signal()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}

	consumer := func() {
		for {
			lock.Lock()
			var v int
			for {
				var ok bool
				if v, ok = queue.Dequeue(); !ok {
					fmt.Println("Queue is empty")
					emptyCond.Wait()
					continue
				}
				break
			}
			lock.Unlock()
			fullCond.Signal()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			fmt.Println(v)
		}
	}
	for i := 0; i < 10; i++ {
		go producer()
	}

	for i := 0; i < 20; i++ {
		go consumer()
	}

	select {} // Wait indefinitely
}
