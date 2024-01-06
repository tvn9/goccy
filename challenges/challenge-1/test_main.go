package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello Another Day!"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	if !strings.Contains(output, msg) {
		t.Errorf("Expected %s, but got %s\n", msg, result)
	}
}

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("Hello Another Day!")
	wg.Wait()
	if msg != "Hello Another Day!" {
		t.Errorf("expected %s, but it is not found\n", msg)
	}
}

func Test_main(t *testing.T) {
	stdOut := out.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("expected %s, but got %s\n", "Hello, universe!", result)
	}
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("expected %s, but got %s\n", "Hello, cosmos!", result)
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("expected %s, but got %s\n", "Hello, world!", result)
	}
}
