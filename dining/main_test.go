package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i <= 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("expected 5, but got %d\n", len(orderFinished))
		}
	}
}

func Test_dineRandTime(t *testing.T) {
	for i := 0; i <= 5; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("expected 5, but got %d\n", len(orderFinished))
		}
	}
}
