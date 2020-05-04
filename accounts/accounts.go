package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string
	balance int
}

// error
var errNoMoney error = errors.New("Can't withdraw")

// Create new Account
func CreateAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit Amount on the Account
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

// Withdraw Amount from the Account
func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errNoMoney
	}

	account.balance -= amount

	return nil
}

// Change Owner of the Account
func (account *Account) ChangeOwner(owner string) {
	account.owner = owner
}

// Get Balance of the Account
func (account Account) GetBalance() int {
	return account.balance
}

// Get Owner of the Account
func (account Account) GetOwner() string {
	return account.owner
}

func (account Account) String() string {
	return fmt.Sprint(account.GetOwner(), "'s account.\nHas: ", account.GetBalance())
}
