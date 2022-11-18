/*
7. Read account no., name, balance amount and amount to be deposited to create and print account
details of 2 persons using struct "bank account".
8. Create account details of 2 persons using struct "bank account". Read account no., name, balance amount
and amount to be deposited and print the updated balance  using  go functions.
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
	accs := [2]BankAccount{inputBankAccount(), inputBankAccount()}

	fmt.Println("\nAccounts Created:")
	fmt.Println(accs[0], accs[1])

	accno := 0
	deposit := 0.0

	fmt.Println("\nEnter details of account to deposit:")
	fmt.Println("Enter account no: ")
	fmt.Scan(&accno)
	fmt.Println("Enter deposit amount: ")
	fmt.Scan(&deposit)

	for _, acc := range accs {
		if acc.accountNo == accno {
			fmt.Println("Account found:", acc)
			acc.balance += deposit
			fmt.Println("New balance: ", acc.balance)
		}
	}
}
