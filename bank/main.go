package main

import (
	"fmt"

	"github.com/jonathas/learning-go/bank/accounts"
)

func structsAndPointers() {
	checkingAccount1 := accounts.CheckingAccount{
		Holder: "Jon",
		AgencyNumber: 589,
		AccountNumber: 123456,
		Balance: 125.50}

	fmt.Println(checkingAccount1)

	checkingAccount2 := accounts.CheckingAccount{"Estela",222,111222,200}

	fmt.Println(checkingAccount2)

	checkingAccount3 := accounts.CheckingAccount{"Estela",222,111222,200}

	// Comparing two structs, returns true
	fmt.Println(checkingAccount2 == checkingAccount3)

	myAccount1 := new(accounts.CheckingAccount)
	myAccount1.Holder = "Jon"
	myAccount1.Balance = 500

	fmt.Println(myAccount1)

	myAccount2 := new(accounts.CheckingAccount)
	myAccount2.Holder = "Jon"
	myAccount2.Balance = 500

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

func main() {
	structsAndPointers()

	silviaAccount := accounts.CheckingAccount{}
	silviaAccount.Holder = "Silvia"
	silviaAccount.Balance = 500

	fmt.Println(silviaAccount)

	fmt.Println(silviaAccount.Withdraw(100.0))

	status, value := silviaAccount.Deposit(300.0)
	fmt.Println(status, value)

	jonAccount := new(accounts.CheckingAccount)
	jonAccount.Holder = "Jon"
	jonAccount.Balance = 500

	fmt.Println("Jon's account balance:", jonAccount.Balance)
	fmt.Println("Silvia's account balance:", silviaAccount.Balance)

	fmt.Println("Silvia transferring 100 to Jon:")
	silviaAccount.Transfer(100, jonAccount)
	fmt.Println("Silvia's account balance after transfer:", silviaAccount.Balance)

	fmt.Println("Jon's account balance after transfer:", jonAccount.Balance)

	fmt.Println(sum(1))
	fmt.Println(sum(1,1))
	fmt.Println(sum(1,1,1))
	fmt.Println(sum(1,1,2,4))
}
