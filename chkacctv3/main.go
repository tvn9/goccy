package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

// Using mutex
type Account struct {
	sync.Mutex
	ID      string
	Balance int
}

// Transfer moves money from one account to another
func Transfer(from, to *Account, amt int) error {
	from.Lock()
	defer from.Unlock()
	to.Lock()
	defer to.Unlock()
	if from.Balance < amt {
		return ErrInsufficient
	}
	from.Balance -= amt
	to.Balance += amt
	return nil
}

func (a *Account) Withdraw(amount int) error {
	a.Lock()
	defer a.Unlock()
	if a.Balance < amount {
		return ErrInsufficient
	}
	a.Balance -= amount
	return nil
}

var ErrInsufficient = errors.New("insufficient fund")

func main() {

	acct1 := Account{Balance: 10}
	acct2 := Account{Balance: 15}
	fmt.Printf("Account 1: %d\n", acct1.Balance)
	fmt.Printf("Account 2: %d\n", acct2.Balance)

	func() {
		fmt.Println("Try to withdraw $5 from Account 1")
		err := acct1.Withdraw(5)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Successfully withdrawn $%d\n", 5)
			fmt.Printf("Account 1 balance: %d\n", acct1.Balance)
		}
	}()
	func() {
		fmt.Println("Try to withdraw $8 from Account 1")
		err := acct1.Withdraw(8)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Withdraw amount: %d\n", 8)
			fmt.Printf("Account 1 balance: %d\n", acct1.Balance)
		}
	}()

	go func() {
		fmt.Printf("Account 1: %d\n", acct1.Balance)
		fmt.Printf("Account 2: %d\n", acct2.Balance)
		fmt.Println("Try to transfer $5 from Account-1 to Account-2")
		err := Transfer(&acct1, &acct2, 5)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("Account 1: %d\n", acct1.Balance)
			fmt.Printf("Account 2: %d\n", acct2.Balance)
		}
	}()
	time.Sleep(1 * time.Second)

}
