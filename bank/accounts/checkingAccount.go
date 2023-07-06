package accounts

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

// c points to who's calling the function
func (c *CheckingAccount) withdraw(value float64) string {
	fmt.Println("Withdrawing", value)

	isWithDrawAllowed := value <= c.balance && value > 0

	if isWithDrawAllowed {
		c.balance -= value
		return "Withdrawal successful"
	}
	return "Insufficient funds"
}

func (c *CheckingAccount) deposit(value float64) (string, float64) {
	fmt.Println("Depositing", value)

	if value > 0 {
		c.balance += value
		return "Deposit successful", c.balance
	}
	return "Invalid deposit value", c.balance
}

func (c *CheckingAccount) transfer(value float64, destinationAccount *CheckingAccount) bool {
	fmt.Println("Transferring", value)

	if value < c.balance && value > 0 {
		c.balance -= value
		destinationAccount.balance += value
		return true
	}
	return false
}
