package main

import (
	"fmt"

	"github.com/jonathas/bank/accounts"
)

func structsAndPointers() {
	checkingAccount1 := accounts.CheckingAccount{
		holder: "Jon",
		agencyNumber: 589,
		accountNumber: 123456,
		balance: 125.50}

	fmt.Println(checkingAccount1)

	checkingAccount2 := accounts.CheckingAccount{"Estela",222,111222,200}

	fmt.Println(checkingAccount2)

	checkingAccount3 := accounts.CheckingAccount{"Estela",222,111222,200}

	// Comparing two structs, returns true
	fmt.Println(checkingAccount2 == checkingAccount3)

	myAccount1 := new(accounts.CheckingAccount)
	myAccount1.holder = "Jon"
	myAccount1.balance = 500

	fmt.Println(myAccount1)

	myAccount2 := new(accounts.CheckingAccount)
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
	silviaAccount.holder = "Silvia"
	silviaAccount.balance = 500

	fmt.Println(silviaAccount)

	fmt.Println(silviaAccount.withdraw(100.0))

	status, value := silviaAccount.deposit(300.0)
	fmt.Println(status, value)

	jonAccount := new(accounts.CheckingAccount)
	jonAccount.holder = "Jon"
	jonAccount.balance = 500

	fmt.Println("Jon's account balance:", jonAccount.balance)
	fmt.Println("Silvia's account balance:", silviaAccount.balance)

	fmt.Println("Silvia transferring 100 to Jon:")
	silviaAccount.transfer(100, jonAccount)
	fmt.Println("Silvia's account balance after transfer:", silviaAccount.balance)

	fmt.Println("Jon's account balance after transfer:", jonAccount.balance)

	fmt.Println(sum(1))
	fmt.Println(sum(1,1))
	fmt.Println(sum(1,1,1))
	fmt.Println(sum(1,1,2,4))
}
