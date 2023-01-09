/*
11. The Fibonacci sequence is defined as: fib (0) = 0, fib (1) = 1,
fib(n) = fib(n-1) + fib(n-2). Write a recursive function in go which can find fib(n).
*/

package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	var n int

	for {
		fmt.Println("\nEnter position of fibonacci number(Starts from 0)(Neg Number to Exit):")
		fmt.Scan(&n)

		if n < 0 {
			break
		}

		fmt.Println("Fibonacci number at position", n, "is", fib(n))
	}
}
