package accounts

import (
	"fmt"

	"github.com/jonathas/learning-go/bank/customers"
)

type SavingsAccount struct {
	Holder 																	customers.Holder
	AgencyNumber, AccountNumber, Operation 	int
	balance 																float64
}

func (c *SavingsAccount) Withdraw(value float64) string {
	fmt.Println("Withdrawing", value)

	isWithDrawAllowed := value <= c.balance && value > 0

	if isWithDrawAllowed {
		c.balance -= value
		return "Withdrawal successful"
	}
	return "Insufficient funds"
}

func (c *SavingsAccount) Deposit(value float64) (string, float64) {
	fmt.Println("Depositing", value)

	if value > 0 {
		c.balance += value
		return "Deposit successful", c.balance
	}
	return "Invalid deposit value", c.balance
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
