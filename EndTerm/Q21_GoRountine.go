/*
Write a go program to implement go routines and channels.

To do this print a Fibonacci series by communicating with main routine through channels.
Also implement a palindrome task to demonstrate concurrent execution with Fibonacci task by showing time consumption of both routines.

[Assume if any other information required]
*/

package main

import (
	"fmt"
	"time"
)

func fibonacciRoutine(n int, ch chan int) {
	start := time.Now()
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	fmt.Println("\nTime taken for fibonacci program: ", time.Since(start))
	close(ch)
}

func palindromeRoutine(s string, ch chan bool) {
	start := time.Now()
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			ch <- false
			return
		}
	}
	fmt.Println("\nTime taken for palindrome program: ", time.Since(start))
	ch <- true
	close(ch)
}

func main() {
	fibonacci := make(chan int)
	palindrome := make(chan bool)

	go fibonacciRoutine(10, fibonacci)
	go palindromeRoutine("madam", palindrome)

	fmt.Println("Fibonacci series: ")
	for i := range fibonacci {
		fmt.Print(i, " ")
	}
	pal := <-palindrome
	fmt.Println("Is 'madam' a palindrome: ", pal)
	fmt.Println("By Karthik Banjan")
}
