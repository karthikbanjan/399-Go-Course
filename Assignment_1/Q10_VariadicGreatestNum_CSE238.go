/*
10.Write a function with one variadic parameter that accepts
a slice, and then it finds the greatest number in a list of numbers.
*/

package main

import "fmt"

func findGreatest(num ...int) {
	max := num[0]

	for _, n := range num {
		if n > max {
			max = n
		}
	}

	fmt.Println("Greatest number is: ", max)
}

func main() {
	list := []int{545, 425, 865, 374, 583, 573, 967}
	findGreatest(list...)
}
