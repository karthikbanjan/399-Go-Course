/*
Given a number n, for each integer i in the
range from 1 to n inclusive, print one value
per line as follows:
• If i is a multiple of both 3 and 5, print
FizzBuzz.
• If i is a multiple of 3 (but not 5), print Fizz.
• If i is a multiple of 5 (but not 3), print Buzz.
• If i is not a multiple of 3 or 5, print the value
of i.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
