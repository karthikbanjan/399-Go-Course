/*
Read an integer number 'N',  write a go program to generate a map with i and square of  i,
such that i is an integer between 1 and N(both included). The program should then print the map
*/

package main

import "fmt"

func main() {
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	squareMap := make(map[int]int)

	for i := 1; i <= N; i++ {
		squareMap[i] = i * i
	}

	fmt.Println(squareMap)
	fmt.Println("Code by Karthik Banjan")
}
