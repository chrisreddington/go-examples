package accounts

import (
	"errors"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []rune("0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

type CurrentAccount struct {
	accountNumber string
	balance       float64
}

func (account CurrentAccount) OpenAccount() IAccount {
	return &CurrentAccount{
		accountNumber: "C-" + RandomString(8),
		balance:       0.00,
	}
}

func (account *CurrentAccount) Deposit(amount float64) error {
	account.balance += amount
	return nil
}

func (account *CurrentAccount) GetAccountNumber() string {
	return account.accountNumber
}

func (account *CurrentAccount) GetBalance() float64 {
	return account.balance
}

func (account *CurrentAccount) Transfer(receiverAccount string, amount float64) error {
	account.balance -= amount
	// TODO: Transfer to receiver account
	return nil
}

func (account *CurrentAccount) Withdraw(amount float64) error {
	account.balance -= amount
	return nil
}

/*
*
* TODO: Implement SavingsAccount
*
 */

type SavingsAccount struct {
	accountNumber string
	created       time.Time
	savings       float64
}

func (account SavingsAccount) OpenAccount() IAccount {
	return &SavingsAccount{
		accountNumber: "S-" + RandomString(8),
		savings:       0.00,
		created:       time.Now(),
	}
}

func (account *SavingsAccount) Deposit(amount float64) error {
	account.savings += amount
	return nil
}

func (account *SavingsAccount) GetAccountNumber() string {
	return account.accountNumber
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.savings
}

func (account SavingsAccount) Transfer(receiverAccount string, amount float64) error {
	return errors.New("You cannot transfer from your savings account")
}

func (account *SavingsAccount) Withdraw(amount float64) error {

	if time.Now().Before(account.created.AddDate(0, 0, 90)) {
		return errors.New("You cannot withdraw from your savings account until 90 days after opening")
	}

	account.savings -= amount
	return nil
}

type ISAAccount struct {
	accountNumber      string
	balance            float64
	remainingAllowance float64
}

func (account ISAAccount) OpenAccount() IAccount {
	return &ISAAccount{
		accountNumber:      "I-" + RandomString(8),
		balance:            0.00,
		remainingAllowance: 400.00,
	}
}

func (account *ISAAccount) Deposit(amount float64) error {

	if (account.remainingAllowance - amount) > 0 {
		account.balance += amount
		account.remainingAllowance -= amount
		return nil
	}

	return errors.New("You cannot deposit more than your remaining allowance")
}

func (account *ISAAccount) GetAccountNumber() string {
	return account.accountNumber
}

func (account *ISAAccount) GetBalance() float64 {
	return account.balance
}

func (account *ISAAccount) Transfer(receiverAccount string, amount float64) error {
	return errors.New("You cannot transfer from your ISA account")
}

func (account *ISAAccount) Withdraw(amount float64) error {

	newBalance := account.balance - (amount + 5.00)

	if newBalance < 0 {
		return errors.New("You cannot withdraw more than your remaining allowance. Make sure to factor in the £5 fee")
	}

	// Take a fee of 5.00 of the currency
	account.balance = newBalance
	return nil
}
