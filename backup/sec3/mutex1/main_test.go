package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	var mutex sync.Mutex

	testMsg := "Goodbye, cruel world!"
	wg.Add(2)
	go updateMessage(testMsg, &mutex)
	go updateMessage("Test 2", &mutex)
	wg.Wait()

	if msg != testMsg {
		t.Errorf("incorrect value, expected %s, but got %s\n", testMsg, msg)
	} else {
		fmt.Println(msg)
	}
}
