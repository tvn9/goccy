package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Account struct {
	sync.Mutex
	ID      string
	Balance int
}

/*
func (a *Account) Withdraw(amount int) error {
	a.Lock()
	defer a.Unlock()
	if a.Balance < amount {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	return nil
}
*/

func Transfer(from, to *Account, amount int) error {
	from.Lock()
	defer from.Unlock()
	to.Lock()
	defer to.Unlock()
	if from.Balance < amount {
		return errors.New("insufficient funds")
	}
	from.Balance -= amount
	to.Balance += amount
	return nil
}

func main() {
	acct1 := Account{
		Balance: 10,
	}
	acct2 := Account{
		Balance: 15,
	}

	go func() {
		if err := Transfer(&acct1, &acct2, 5); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		if err := Transfer(&acct2, &acct1, 5); err != nil {
			fmt.Println(err)
		}
	}()
	time.Sleep(1 * time.Second)
}
