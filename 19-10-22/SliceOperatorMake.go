package main

import "fmt"

func main() {
	x := make([]int, 5, 10)

	fmt.Println("Enter the values: ")
	for i := range x {
		fmt.Scanln(&x[i])
	}

	fmt.Println("The values are: ")
	for _, v := range x {
		fmt.Println(v)
	}
}
