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
	acct := Account{
		Balance: 10,
	}
	fmt.Println("Current account balance:", acct.Balance)
	go func() {
		fmt.Println("Try to withdraw $5")
		err := acct.Withdraw(5)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("Withdraw amount %d\n", 5)
			fmt.Printf("Balance: %d\n", acct.Balance)
		}
	}()
	go func() {
		fmt.Println("Try to withdraw 8")
		err := acct.Withdraw(8)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("Withdraw amount: %d\n", 8)
			fmt.Printf("Balance: %d\n", acct.Balance)
		}
	}()
	time.Sleep(1 * time.Second)
}
