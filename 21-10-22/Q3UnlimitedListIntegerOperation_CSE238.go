//By Karthik Banjan

package main

import "fmt"

func main() {
	a := 0
	var input []int

	fmt.Print("Enter numbers: ")
	for {
		fmt.Scan(&a)
		if a == 0 {
			break
		}
		input = append(input, a)
	}

	var square []int
	var positive []int

	for _, v := range input {
		square = append(square, v*v)
		if v > 0 {
			positive = append(positive, v)
		}
	}

	fmt.Println("Square: ", square)
	fmt.Println("Positive: ", positive)
}
