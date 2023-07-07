package main

import (
	"fmt"

	"github.com/jonathas/learning-go/bank/accounts"
	"github.com/jonathas/learning-go/bank/customers"
)

func structsAndPointers() {
	jon := customers.Holder{Name:"Jon", Document:"123.123.123-12", Job:"Developer"}

	checkingAccount1 := accounts.CheckingAccount{
		Holder: jon,
		AgencyNumber: 589,
		AccountNumber: 123456}

	checkingAccount1.Deposit(125.50)

	fmt.Println(checkingAccount1)

	estela := customers.Holder{Name:"Estela"}
	checkingAccount2 := accounts.CheckingAccount{Holder:estela,AgencyNumber:222,AccountNumber:111222}
	checkingAccount2.Deposit(200)

	fmt.Println(checkingAccount2)

	checkingAccount3 := accounts.CheckingAccount{Holder:estela,AgencyNumber:222,AccountNumber:111222}
	checkingAccount3.Deposit(200)

	// Comparing two structs, returns true
	fmt.Println(checkingAccount2 == checkingAccount3)

	myAccount1 := new(accounts.CheckingAccount)
	myAccount1.Holder = jon

	myAccount1.Deposit(500)

	fmt.Println(myAccount1)

	myAccount2 := new(accounts.CheckingAccount)
	myAccount2.Holder = jon

	myAccount2.Deposit(500)

	// Comparing two pointers, returns false
	fmt.Println(myAccount1 == myAccount2)

	// Because the memory addresses are different, even though the content is the same
	fmt.Println(&myAccount1)
	fmt.Println(&myAccount2)

	// However, if we compare the content of the pointers, it returns true
	fmt.Println(*myAccount1 == *myAccount2)
}

// Variadic function, can receive any number of parameters
func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
			result += number
	}
	return result
}

func PayInvoice(account verifyAccount, value float64) {
	account.Withdraw(value)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	structsAndPointers()

	silvia := customers.Holder{Name:"Silvia"}
	jon := customers.Holder{Name:"Jon"}

	silviaAccount := accounts.CheckingAccount{}
	silviaAccount.Holder = silvia

	silviaAccount.Deposit(500)

	fmt.Println(silviaAccount)

	fmt.Println(silviaAccount.Withdraw(100.0))

	status, value := silviaAccount.Deposit(300.0)
	fmt.Println(status, value)

	jonAccount := new(accounts.CheckingAccount)
	jonAccount.Holder = jon

	jonAccount.Deposit(500)

	fmt.Println("Jon's account balance:", jonAccount.GetBalance())
	fmt.Println("Silvia's account balance:", silviaAccount.GetBalance())

	fmt.Println("Silvia transferring 100 to Jon:")
	silviaAccount.Transfer(100, jonAccount)
	fmt.Println("Silvia's account balance after transfer:", silviaAccount.GetBalance())

	fmt.Println("Jon's account balance after transfer:", jonAccount.GetBalance())

	fmt.Println(sum(1))
	fmt.Println(sum(1,1))
	fmt.Println(sum(1,1,1))
	fmt.Println(sum(1,1,2,4))

	/**
	* Doesn't matter if the account is checking or savings, as long as it implements the Withdraw method,
	* it can be used in the PayInvoice function with the verifyAccount interface
	*/
	denisAccount := accounts.SavingsAccount{}
	denisAccount.Deposit(500)
	fmt.Println(denisAccount.GetBalance())

	PayInvoice(&denisAccount, 100)
	fmt.Println(denisAccount.GetBalance())

	maryAccount := accounts.CheckingAccount{}
	maryAccount.Deposit(500)
	fmt.Println(maryAccount.GetBalance())

	PayInvoice(&maryAccount, 300)
	fmt.Println(maryAccount.GetBalance())
}
