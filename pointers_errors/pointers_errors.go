package pointers_errors

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("can not withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

//type Cart struct {
//	iphone int
//	laptop int
//}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	if w.balance < amount {
		return ErrorInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//func (c *Cart) CartDeposit(amount int) {
//	fmt.Printf("address of cart-iphone  is %p and cart-laptop %p in deposit \n", &c.iphone, &c.laptop)
//	c.iphone += amount
//	c.laptop += amount
//}
//
//func (c *Cart) CartBalance() int {
//	return c.laptop + c.iphone
//}
