package main

import "fmt"

func main() {
	var x float64
	var y float64
	var choice int

	fmt.Println("Hello Welcome to the Calculator")

	for {
		fmt.Println("Enter the first number")
		fmt.Scanln(&x)
		fmt.Println("Enter the second number")
		fmt.Scanln(&y)

		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Addition of two numbers is", x+y)

		case 2:
			fmt.Println("Subtraction of two numbers is", x-y)

		case 3:
			fmt.Println("Multiplication of two numbers is", x*y)

		case 4:
			fmt.Println("Division of two numbers is", x/y)

		case 5:
			fmt.Println("Thank you for using the calculator")

		default:
			fmt.Println("Invalid choice...Restarting!")
		}
	}
}
