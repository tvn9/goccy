package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printing(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)
	go printing("Thanh Nguyen", &wg)
	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Thanh Nguyen") {
		t.Errorf("expected %s, but got %s instead", result, output)
	} else {
		fmt.Printf("Pass - expected %s, and got %s", result, output)
	}
}
