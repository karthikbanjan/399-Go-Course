/*
Create a struct called products, which describe products in a super market.
Read code number, price, item type and quantity of each unit.
Make a function "details" and initialize any 2 products.
Make another function to update product details(only if user wants to update).
*/

package main

import "fmt"

type products struct {
	codeNumber int
	price      float64
	itemType   string
	quantity   int
}

var p1 products
var p2 products

func details() {
	fmt.Println("Please enter the details of the product 1:")

	fmt.Print("Code Number: ")
	fmt.Scan(&p1.codeNumber)
	fmt.Print("Price: ")
	fmt.Scan(&p1.price)
	fmt.Print("Item Type: ")
	fmt.Scan(&p1.itemType)
	fmt.Print("Quantity: ")
	fmt.Scan(&p1.quantity)

	fmt.Println("Please enter the details of the product 2:")
	fmt.Print("Code Number: ")
	fmt.Scan(&p2.codeNumber)
	fmt.Print("Price: ")
	fmt.Scan(&p2.price)
	fmt.Print("Item Type: ")
	fmt.Scan(&p2.itemType)
	fmt.Print("Quantity: ")
	fmt.Scan(&p2.quantity)
}

func update(choice2 int) {
	fmt.Println("Please enter the details of the product:")

	if choice2 == 1 {
		fmt.Print("Code Number: ")
		fmt.Scan(&p1.codeNumber)
		fmt.Print("Price: ")
		fmt.Scan(&p1.price)
		fmt.Print("Item Type: ")
		fmt.Scan(&p1.itemType)
		fmt.Print("Quantity: ")
		fmt.Scan(&p1.quantity)
		fmt.Println("Product 1: ", p1)
	} else {
		fmt.Print("Code Number: ")
		fmt.Scan(&p2.codeNumber)
		fmt.Print("Price: ")
		fmt.Scan(&p2.price)
		fmt.Print("Item Type: ")
		fmt.Scan(&p2.itemType)
		fmt.Print("Quantity: ")
		fmt.Scan(&p2.quantity)
		fmt.Println("Product 2: ", p2)
	}
}

func main() {
	details()

	fmt.Println("Product 1: ", p1)
	fmt.Println("Product 2: ", p2)

	fmt.Println("Do you wish to update the details of any product? (Y/N)")
	var choice string
	fmt.Scan(&choice)
	if choice == "Y" {
		fmt.Println("Which product do you wish to update? (1/2)")
		var choice2 int
		fmt.Scan(&choice2)
		update(choice2)
	}
	fmt.Println("Code by Karthik Banjan")
}
