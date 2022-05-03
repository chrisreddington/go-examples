package accounts

type IAccount interface {
	Deposit(amount float64) error
	GetAccountNumber() string
	GetBalance() float64
	OpenAccount() IAccount
	Transfer(receiverAccount string, amount float64) error
	Withdraw(amount float64) error
}
