/*
Write a go application to calculate electricity bill.
Accept number of units used from user. If no. of units used is less than 40, then every unit costs Rs.7.

If units used  is above 40 and less than 100, then additional units from 41 to 100 should pay Rs.15 per unit.
If units used is above 100, then entire unit should pay Rs.20 per unit.
*/

package main

import "fmt"

func main() {
	var units int
	println("Enter units: ")
	fmt.Scan(&units)

	var bill float64
	var extraUnits float64

	if units > 100 {
		bill = float64(units) * 20.0
	} else if units > 40 {
		extraUnits = float64(units - 40)
		bill = (40 * 7.0) + (extraUnits * 15.0)
	} else {
		bill = float64(units) * 7.0
	}

	fmt.Println("Bill: ", bill)
	fmt.Println("Code by Karthik Banjan")
}
