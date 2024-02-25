package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	weekly := "Weekly total income:          660.00"
	yearly := "52 weeks - Total Balance:   $34320.00"

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, weekly) {
		t.Errorf("expected %s, but got %s\n", weekly, output)
	} else {
		fmt.Printf("Pass - expected %s, and got %s\n", weekly, output)
	}

	if !strings.Contains(output, yearly) {
		t.Errorf("expected %s, but got %s\n", yearly, output)
	} else {
		fmt.Printf("Pass - expected %s, and got %s\n", yearly, output)
	}

}
