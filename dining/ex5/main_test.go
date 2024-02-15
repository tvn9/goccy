package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_dine(t *testing.T) {

	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 20; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d\n", len(orderFinished))
		} else {
			fmt.Println("Test Passed, expected 5 but got", len(orderFinished))
		}
	}
}

func Test_dineWithDelay(t *testing.T) {
	// test table
	test_delay := []struct {
		name     string
		duration time.Duration
	}{
		{"No delay", (0 * time.Second)},
		{"with delay 100", time.Millisecond * 100},
		{"With delay 400", time.Millisecond * 400},
	}

	for _, test := range test_delay {
		orderFinished = []string{}

		eatTime = test.duration
		sleepTime = test.duration
		thinkTime = test.duration

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice; expected 5 but got %d\n", test.name,
				len(orderFinished))
		} else {
			fmt.Printf("%s passed; expected 5, got %d\n", test.name, len(orderFinished))
		}
	}
}
