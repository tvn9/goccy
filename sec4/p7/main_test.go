package main

import (
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	delay = 0 * time.Second

	for i := 0; i < 10000; i++ {
		main()
		if len(orderFinished) != 5 {
			t.Error("not everyone finish eating!.")
		}
		orderFinished = []string{}
	}
}
