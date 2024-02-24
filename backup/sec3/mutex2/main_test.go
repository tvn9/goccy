package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_MainJobIncome(t *testing.T) {
	// Testging struct fields and total bankBalance
	var bankBalance int

	incomes := Income{}

	incomes.Source = "Main Job"
	incomes.Amount = 500
	bankBalance += incomes.Amount

	if bankBalance != 500 {
		t.Errorf("expected bank balance %d.00 but got %d.00", 500, bankBalance)
	}

	incomes.Source = "Gifts"
	incomes.Amount = 20
	bankBalance += incomes.Amount
	if bankBalance != 520 {
		t.Errorf("expected bank balance %d.00, but got %d.00", 520, bankBalance)
	}

	incomes.Source = "Part time job"
	incomes.Amount = 100
	bankBalance += incomes.Amount
	if bankBalance != 620 {
		t.Errorf("expected bank balance %d.00, but got %d.00", 620, bankBalance)
	}

	incomes.Source = "Investments"
	incomes.Amount = 100
	bankBalance += incomes.Amount
	if bankBalance != 720 {
		t.Errorf("expected bank balance %d.00, but got %d.00", 720, bankBalance)
	}

	// Testing main function
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Errorf("Wrong balance returned!")
	}
}
