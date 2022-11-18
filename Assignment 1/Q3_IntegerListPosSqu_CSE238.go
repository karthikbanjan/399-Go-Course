/*
3. Read a list of integers from the user, name it as input, until the user enters 0. Then Create 2 new
lists namely
1). Square, where each element is a square of the corresponding element in the input list
2). Positive, with only positive numbers from the input list.

*/

package main

import "fmt"

func main() {
	var input []int
	var num int

	fmt.Println("Enter numbers (0 to stop):")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		input = append(input, num)
	}

	var squareList []int
	var posList []int

	for _, v := range input {
		squareList = append(squareList, v*v)
		if v > 0 {
			posList = append(posList, v)
		}
	}

	fmt.Println("Squares:", squareList)
	fmt.Println("Positive Numbers:", posList)
}
