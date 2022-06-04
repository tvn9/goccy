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
	if !strings.Contains(output, "Hello Another Day!") {
		t.Errorf("expected %s, but go %s\n", result, output)
	}
}

func Test_updateMesage(t *testing.T) {
	wg.Add(1)

	go updateMessage("Hello Another Day!")

	wg.Wait()

	if msg != "Hello Another Day!" {
		t.Errorf("expected %s, but it is not found\n", msg)
	}
}
