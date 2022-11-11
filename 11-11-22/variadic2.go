/*
Initialize a slice to store float values which is
reading from user. Make a function to pass these values
to a variadic parameter. Finally, the function will
return sum of squares of numbers.
*/

package main

import "fmt"

func sumOfSquares(n ...float64) float64 {
	sum := 0.0
	for _, v := range n {
		sum += v * v
	}
	return sum
}

func main() {
	var n []float64
	var num float64

	for {
		fmt.Print("Enter a number: ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		n = append(n, num)
	}

	fmt.Println("Sum of squares of numbers is", sumOfSquares(n...))
}
