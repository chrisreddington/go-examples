package main

import (
	"fmt"
	"strconv"

	"github.com/chrisreddington/go-examples/interfaces/accounts"
)

var list []accounts.IAccount = []accounts.IAccount{
	&accounts.CurrentAccount{},
	&accounts.SavingsAccount{},
	&accounts.ISAAccount{},
}

func main() {
	for _, account := range list {
		account := account.OpenAccount()

		err := account.Deposit(500.00)
		if err != nil {
			fmt.Println(err)
		}

		err = account.Withdraw(50.00)
		if err != nil {
			fmt.Println(err)
		}

		err = account.Transfer("X-123456", 100.00)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(account.GetAccountNumber() + ": " + strconv.FormatFloat(account.GetBalance(), 'f', 2, 64))
	}
}
