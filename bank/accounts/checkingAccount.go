package accounts

import (
	"fmt"

	"github.com/jonathas/learning-go/bank/customers"
)

/**
 * Attribute names start with uppercase so they can be exported.
 * If they start with lowercase, they can only be used inside the package (private)
 */
type CheckingAccount struct {
	Holder        customers.Holder
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

// c points to who's calling the function
func (c *CheckingAccount) Withdraw(value float64) string {
	fmt.Println("Withdrawing", value)

	isWithDrawAllowed := value <= c.Balance && value > 0

	if isWithDrawAllowed {
		c.Balance -= value
		return "Withdrawal successful"
	}
	return "Insufficient funds"
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	fmt.Println("Depositing", value)

	if value > 0 {
		c.Balance += value
		return "Deposit successful", c.Balance
	}
	return "Invalid deposit value", c.Balance
}

func (c *CheckingAccount) Transfer(value float64, destinationAccount *CheckingAccount) bool {
	fmt.Println("Transferring", value)

	if value < c.Balance && value > 0 {
		c.Balance -= value
		destinationAccount.Balance += value
		return true
	}
	return false
}
