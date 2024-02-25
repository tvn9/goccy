package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// veriable for bank balance
	bankBalance := 0
	weekly := 0

	// print out starting values
	fmt.Printf("Initial account balance $%d.00\n", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	// print out weekly revenue
	for i, income := range incomes {
		wg.Add(1)
		go func(i int, income Income) {
			defer wg.Done()
			fmt.Printf("#%d - %-24s %3d.00\n", i, income.Source, income.Amount)
			weekly += income.Amount
		}(i, income)
		wg.Wait()
	}
	fmt.Printf("Weekly total income: %12d.00\n", weekly)

	// loop through 52 weeks and print out total income
	fmt.Println()
	for i := 1; i <= 52; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bankBalance += weekly
		}()
		wg.Wait()
		fmt.Printf("week %d - income %17d.00\n", i, bankBalance)
	}

	// print out final balance
	fmt.Printf("52 weeks - Total Balance:   $%d.00\n", bankBalance)
}
