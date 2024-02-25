package main

import (
	"fmt"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, World!"
	testMsg := "Hello, Thanh Nguyen!"
	fmt.Printf("message before update: %s\n", msg)
	wg.Add(2)
	go updateMessage(testMsg)
	go updateMessage("Hello, GoPher!")
	wg.Wait()

	if testMsg != msg {
		t.Errorf("expected %s, but got %s\n", testMsg, msg)
	} else {
		fmt.Printf("Pass - expected %s, and got %s\n", testMsg, msg)
	}
}
