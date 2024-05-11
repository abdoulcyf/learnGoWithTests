package more_pointers_errors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type GBP int

type Account struct {
	balance GBP
}

type Stringer interface {
	String() string
}

func (b GBP) String() string {
	return fmt.Sprintf("%d GBP", b)
}

//var ErrorInsufficientFunds = errors.New("insufficient funds")

func (a *Account) Deposit(amount GBP) {
	a.balance += amount
}

func (a *Account) Balance() GBP {
	return a.balance
}

func (a *Account) Withdraw(amount GBP) {
	//
	//if amount > a.balance {
	//	return errors.Unwrap()
	//}

	a.balance -= amount
}
