// Program to project how much money one will make in the future time frame.
package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {

	// define a wait group

	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)

	// define weely revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	// loop through 52 weeks and print out how much is made; keep a running total
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
				balance.Unlock()
			}
		}(i, income)
	}
	wg.Wait()

	// print out final balance
	fmt.Printf("Total account balance: $%d.00\n", bankBalance)
}
