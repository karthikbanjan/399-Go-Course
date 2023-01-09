package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func palindrome(s string) {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			println(s, "is not a palindrome")
			return
		}
	}
	fmt.Println(s, "is a palindrome")
}

func factorial(n int) {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	fmt.Println(n, "! =", f)
}

func main() {
	now := time.Now()

	wg.Add(2)
	go func() {
		defer wg.Done()
		palindrome("kayak")
	}()
	go func() {
		defer wg.Done()
		factorial(5)
	}()
	wg.Wait()

	/*
		palindrome("kayak")
		factorial(5)
	*/
	fmt.Println("Time elapsed:", time.Since(now))
}
