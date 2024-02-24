package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Balance int
}

func (a *Account) Withdraw(amount int) error {
	if a.Balance < amount {
		fmt.Println("If", a.Balance)
		return ErrInsufficient
	}
	a.Balance -= amount
	fmt.Println("else", a.Balance)
	return nil
}

var ErrInsufficient = errors.New("insufficient fund")

func main() {
	acct := Account{
		Balance: 10,
	}
	go fmt.Println("Run 1", acct.Withdraw(6))
	go fmt.Println("Run 2", acct.Withdraw(7))
}
