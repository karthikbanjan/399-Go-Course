/*
A data analytics company collecting body weight of people living in a society.
Write go program to read these values and store in a list, make this process with
an infinite loop, if user enters 0, STOP the process.
Iterate given list and create 2 new lists namely

1). Average, where each weight should be between 45 and 70.

2). Heavy, weight is greater than 70 and less than 110.
if it is more than 110, display 'wrong value'.
*/

package main

import "fmt"

func main() {
	var weight []float64
	var avg []float64
	var heavy []float64

	for {
		var w float64
		fmt.Print("Enter weight: ")
		fmt.Scan(&w)

		if w == 0 {
			break
		}

		weight = append(weight, w)
	}

	for _, v := range weight {
		if v >= 45 && v <= 70 {
			avg = append(avg, v)
		} else if v > 70 && v < 110 {
			heavy = append(heavy, v)
		} else if v > 110 {
			fmt.Println("Wrong value: ", v)
		}
	}

	fmt.Println("Average: ", avg)
	fmt.Println("Heavy: ", heavy)
	fmt.Println("Code by Karthik Banjan")
}
