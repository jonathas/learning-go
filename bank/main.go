package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func structsAndPointers() {
	checkingAccount1 := CheckingAccount{
		holder: "Jon",
		agencyNumber: 589,
		accountNumber: 123456,
		balance: 125.50}

	fmt.Println(checkingAccount1)

	checkingAccount2 := CheckingAccount{"Estela",222,111222,200}

	fmt.Println(checkingAccount2)

	checkingAccount3 := CheckingAccount{"Estela",222,111222,200}

	// Comparing two structs, returns true
	fmt.Println(checkingAccount2 == checkingAccount3)

	myAccount1 := new(CheckingAccount)
	myAccount1.holder = "Jon"
	myAccount1.balance = 500

	fmt.Println(myAccount1)

	myAccount2 := new(CheckingAccount)
	myAccount2.holder = "Jon"
	myAccount2.balance = 500

	// Comparing two pointers, returns false
	fmt.Println(myAccount1 == myAccount2)

	// Because the memory addresses are different, even though the content is the same
	fmt.Println(&myAccount1)
	fmt.Println(&myAccount2)

	// However, if we compare the content of the pointers, it returns true
	fmt.Println(*myAccount1 == *myAccount2)
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

// Variadic function, can receive any number of parameters
func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
			result += number
	}
	return result
}

func main() {
	structsAndPointers()

	silviaAccount := CheckingAccount{}
	silviaAccount.holder = "Silvia"
	silviaAccount.balance = 500

	fmt.Println(silviaAccount)

	fmt.Println(silviaAccount.withdraw(100.0))

	status, value := silviaAccount.deposit(300.0)
	fmt.Println(status, value)

	fmt.Println(sum(1))
	fmt.Println(sum(1,1))
	fmt.Println(sum(1,1,1))
	fmt.Println(sum(1,1,2,4))
}
