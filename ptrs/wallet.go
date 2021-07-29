package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Withdraw(amt Bitcoin) error {

	if amt > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amt
	return nil
}
