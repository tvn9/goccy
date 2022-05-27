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

	msg = "Hello, cosmos!"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("expected %s, but go %s\n", result, output)
	}
}

func Test_updateMesage(t *testing.T) {
	wg.Add(1)

	go updateMessage("Hello, cosmos!")

	wg.Wait()

	if msg != "Hello, cosmos!" {
		t.Errorf("expected %s, but it is not found\n", msg)
	}
}
