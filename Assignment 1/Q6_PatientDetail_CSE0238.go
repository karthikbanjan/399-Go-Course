/*
6. Given a map with covid patient details from a hospital with patient id and age.
data {1001:21, 1002:35, 1003:12, 1004:64, 1005:17, 1006:59, 1007:43........}
Make 2 list, one contains id of young patients less than 18 years old and second one consists of id of
senior citizens aged above 60

*/

package main

import "fmt"

func main() {
	patientDetails := map[int]int{1001: 21, 1002: 35, 1003: 12, 1004: 64,
		1005: 17, 1006: 59, 1007: 43, 1008: 75, 1009: 18, 1010: 60}
	var youngPatients []int
	var seniorPatients []int

	for k, v := range patientDetails {
		if v < 18 {
			youngPatients = append(youngPatients, k)
		} else if v > 60 {
			seniorPatients = append(seniorPatients, k)
		}
	}

	fmt.Println("Young Patients: ", youngPatients)
	fmt.Println("Senior Citizens: ", seniorPatients)
}
