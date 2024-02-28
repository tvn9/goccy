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

func Test_dineTimeTable(t *testing.T) {
	var randTime = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, test := range randTime {
		orderFinished = []string{}

		eatTime = test.delay
		sleepTime = test.delay
		thinkTime = test.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s, expected 5, but got %d\n", test.name, len(orderFinished))
		}
	}
}
