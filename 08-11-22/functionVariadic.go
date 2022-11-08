package main

import "fmt"

func addition(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	var z = []int{1, 2, 3, 4, 5}
	fmt.Println("Sum:", addition(z...))
}
