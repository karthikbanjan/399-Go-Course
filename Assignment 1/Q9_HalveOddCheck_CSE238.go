/*
9. Write a function which takes an integer and halves it and returns true if it is even or false if it is odd.
For example, half (1) should return (0, false) and half (2) should return (1, true).
*/

package main

import "fmt"

func halveCheck(num int) (int, bool) {
	return num / 2, num%2 == 0
}

func main() {
	var num int
	fmt.Println("Enter a integer: ")
	fmt.Scan(&num)

	half, isEven := halveCheck(num)

	fmt.Println("Half of", num, "is", half)
	fmt.Println("Is", num, "even?", isEven)
}
