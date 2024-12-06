package wallet

import "errors"

type Wallet struct {
	Balance int
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.Balance += amount
}

func (wallet *Wallet) Withdraw(amount int) error {
	if amount > wallet.Balance {
		return errors.New("not enough money, please try again")
	}

	wallet.Balance -= amount
	return nil
}
