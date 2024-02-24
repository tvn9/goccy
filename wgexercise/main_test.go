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
	fmt.Println("Orginal message:", msg)
	testMsg := "Thanh Nguyen"
	wg.Add(1)
	go updateMessage(testMsg, &wg)
	wg.Wait()
	fmt.Print("Updated message: ")
	printMessage()

	if testMsg != msg {
		t.Errorf("expected %s, but got %s\n", msg, msg)
	} else {
		fmt.Printf("Pass - expected %s and got %s\n", msg, msg)
	}
}

func Test_printMessage(t *testing.T) {
	msg = "Hello, World!"
	testMsg := "Thanh Nguyen"
	// redirecting standard output to a variable
	stdOut := os.Stdout

	// redirecting pipe file io to r (read), w (write), ignore error
	r, w, _ := os.Pipe()
	os.Stdout = w

	wg.Add(1)
	updateMessage(testMsg, &wg)
	wg.Wait()
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, testMsg) {
		t.Errorf("expect %s, but got %s\n", result, output)
	} else {
		fmt.Printf("expected %s, and got %s\n", result, output)
	}

}
