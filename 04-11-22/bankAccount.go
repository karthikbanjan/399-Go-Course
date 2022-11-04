package main

import "fmt"

type BankAccount struct {
	accountNo int
	owner     string
	balance   float64
}

func main() {
	bankAccount1 := BankAccount{accountNo: 1234, owner: "John Doe", balance: 1000.0}
	bankAccount2 := BankAccount{accountNo: 5678, owner: "Jane Doe", balance: 2000.0}

	an := 0
	deposit := 0
	fmt.Println("Enter account number: ")
	fmt.Scan(&an)
	fmt.Println("Enter deposit amount: ")
	fmt.Scan(&deposit)

	if bankAccount1.accountNo == an {
		bankAccount1.balance += float64(deposit)
		fmt.Println("New balance: ", bankAccount1.balance)
	} else if bankAccount2.accountNo == an {
		bankAccount2.balance += float64(deposit)
		fmt.Println("New balance: ", bankAccount2.balance)
	} else {
		fmt.Println("Invalid account number")
	}

	fmt.Println(bankAccount1)
	fmt.Println(bankAccount2)
}
