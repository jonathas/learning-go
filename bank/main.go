package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
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
