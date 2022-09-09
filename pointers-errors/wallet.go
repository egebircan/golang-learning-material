package pointers_errors

import (
	"errors"
	"fmt"
)

// It seems a bit overkill to create a struct for this. int is fine in terms of the way it works but it's not descriptive.
// Go lets you create new types from existing ones.
// By doing this we're making a new type and we can declare methods on them.
// This can be very useful when you want to add some domain specific functionality on top of existing types.
// Lets you implement interfaces
type Bitcoin int

// could be omitted because this is already defined in fmt package
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// In Go, when you call a function or a method the arguments are copied.
// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.
// So we should use pointers instead
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// In Go, errors are values
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Technically you do not need to change Balance to use a pointer receiver as taking a copy of the balance is fine.
// However, by convention you should keep your method receiver types the same for consistency.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}