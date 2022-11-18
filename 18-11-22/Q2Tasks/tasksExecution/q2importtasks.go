package main

import (
	"fmt"
	"tasks"
)

func main() {
	var numbers []int
	var choice int

	fmt.Println("Enter numbers (-1 to stop): ")
	for {
		var num int
		fmt.Scan(&num)
		if num == -1 {
			break
		}
		numbers = append(numbers, num)
	}

	fmt.Println("\nEntered numbers: ", numbers)
	fmt.Println()

	for {
		fmt.Println("Tasks:")
		fmt.Println("1. Sum of numbers")
		fmt.Println("2. Sum of squares")
		fmt.Println("3. Sort in ascending order")
		fmt.Println("4. Sort in descending order")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		fmt.Println()

		switch choice {
		case 1:
			fmt.Println("Sum of numbers:", tasks.SumOfNumbers(numbers...))

		case 2:
			fmt.Println("Sum of squares:", tasks.SumOfSquares(numbers...))

		case 3:
			fmt.Println("Sorted in ascending order:", tasks.SortAscending(numbers...))

		case 4:
			fmt.Println("Sorted in descending order:", tasks.SortDescending(numbers...))

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice...Restarting!")
		}
		fmt.Println()
	}
}
