package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	msg = "Hello, World!"
	testMsg := "Thanh Nguyen"

	wg.Add(1)
	updateMessage(testMsg, &wg)
	wg.Wait()

	if msg != testMsg {
		t.Errorf("expect %s, but go %s\n", testMsg, msg)
	} else {
		fmt.Printf("expected %s, and got %s\n", testMsg, msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	msg = "Hello, Thanh Nguyen!"

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, Thanh Nguyen!") {
		t.Errorf("expect %s, but got %s\n", output, msg)
	} else {
		fmt.Printf("expect %s, and got %s\n", output, msg)
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("expect %s, but got %s\n", output, msg)
	} else {
		fmt.Printf("expect %s, and got %s\n", output, msg)
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("expect %s, but got %s\n", output, msg)
	} else {
		fmt.Printf("expect %s, and got %s\n", output, msg)
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("expect %s, but got %s\n", output, msg)
	} else {
		fmt.Printf("expect %s, and got %s\n", output, msg)
	}
}
