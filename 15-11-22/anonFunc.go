/*
anonymous func for multiplication
*/

package main

import "fmt"

var mult = func(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println(mult(5, 6))
}
