/*
Function to swap two integers.
Parameters using pointers
*/

package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	a := 0
	b := 0

	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	fmt.Print("Enter another number: ")
	fmt.Scan(&b)

	swap(&a, &b)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
