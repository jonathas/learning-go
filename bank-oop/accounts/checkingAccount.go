package accounts

import (
	"fmt"

	"github.com/jonathas/learning-go/bank-oop/customers"
)

/**
 * Attribute names start with uppercase so they can be exported.
 * If they start with lowercase, they can only be used inside the package (private)
 */
type CheckingAccount struct {
	Holder        customers.Holder
	AgencyNumber  int
	AccountNumber int
	balance       float64 // private attribute (encapsulated!)
}

// c points to who's calling the function
func (c *CheckingAccount) Withdraw(value float64) string {
	fmt.Println("Withdrawing", value)

	isWithDrawAllowed := value <= c.balance && value > 0

	if isWithDrawAllowed {
		c.balance -= value
		return "Withdrawal successful"
	}
	return "Insufficient funds"
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	fmt.Println("Depositing", value)

	if value > 0 {
		c.balance += value
		return "Deposit successful", c.balance
	}
	return "Invalid deposit value", c.balance
}

func (c *CheckingAccount) Transfer(value float64, destinationAccount *CheckingAccount) bool {
	fmt.Println("Transferring", value)

	if value < c.balance && value > 0 {
		c.balance -= value
		destinationAccount.balance += value
		return true
	}
	return false
}

// in order to access the balance attribute, we need to create a getter
func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
