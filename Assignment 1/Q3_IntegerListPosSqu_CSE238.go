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
