package main

// Simple cache implementation. This cache may load the schema element multiple times.
// You may get a different number of cache misses every time you run it.

import (
	"fmt"
	"math/rand"
	"sync"
)

type Data struct {
	ID      string
	Payload string
}

type Cache struct {
	mu sync.Mutex
	m  map[string]*Data
}

var (
	numCalls     int
	numCallsLock sync.Mutex
)

func retrieveData(ID string) (*Data, bool) {
	numCallsLock.Lock()
	defer numCallsLock.Unlock()
	numCalls++
	return &Data{ID: ID, Payload: "payload"}, true
}

func (c *Cache) Get(ID string) (Data, bool) {
	c.mu.Lock()
	data, ok := c.m[ID]
	c.mu.Unlock()
	if ok {
		if data == nil {
			return Data{}, false
		}
		return *data, true
	}
	data, loaded := retrieveData(ID)
	c.mu.Lock()
	defer c.mu.Unlock()
	d, exists := c.m[data.ID]
	if exists {
		return *d, true
	}
	if !loaded {
		c.m[ID] = nil
		return Data{}, false
	}
	c.m[data.ID] = data
	return *data, true
}

func main() {
	cache := Cache{m: make(map[string]*Data)}

	ids := []string{"a", "b", "c", "e", "f", "g", "h", "i", "j"}

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				id := ids[rand.Intn(len(ids))]
				cache.Get(id)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Number of cache misses: %d\n", numCalls)
}
