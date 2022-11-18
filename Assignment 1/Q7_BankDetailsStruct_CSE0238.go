/*
7. Read account no., name, balance amount and amount to be deposited to create and print account
details of 2 persons using struct "bank account".
*/

package main

import "fmt"

type BankAccount struct {
	accountNo int
	ownerName string
	balance   float64
}

func inputBankAccount() BankAccount {
	var account BankAccount
	fmt.Println("Enter account details:")
	fmt.Println("Enter account no: ")
	fmt.Scan(&account.accountNo)

	var firstName, lastName string
	fmt.Println("Enter owner's first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter owner's last name:")
	fmt.Scan(&lastName)
	account.ownerName = firstName + " " + lastName

	fmt.Println("Enter balance: ")
	fmt.Scan(&account.balance)
	return account
}

func main() {
	acc1 := inputBankAccount()
	acc2 := inputBankAccount()

	fmt.Println(acc1)
	fmt.Println(acc2)
}
