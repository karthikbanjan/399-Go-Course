package main

import "fmt"

func square(x *int) {
	*x = *x * *x
}

func main() {
	var in int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &in)
	square(&in)
	fmt.Println(in)
}
