package main

import (
	"fmt"
)

func main() {
	data := map[int]int{1001: 21, 1002: 35, 1003: 12, 1004: 64, 1005: 17, 1006: 59, 1007: 43}

	for true {
		fmt.Println()

		fmt.Println("Enter the patient ID:")
		var patID int
		fmt.Scanln(&patID)

		fmt.Println("Covid Details Updation Menu:")
		fmt.Println("1. Vaccinated")
		fmt.Println("2. Not Vaccinated")
		fmt.Println("3. Exit")
		fmt.Println("Enter your choice:")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			delete(data, patID)
			fmt.Println("Updated Map:", data)

		case 2:
			data[patID]++
			fmt.Println("Updated Map:", data)
			var notVaccinated []int
			for k, _ := range data {
				notVaccinated = append(notVaccinated, k)
			}

			fmt.Println("Not Vaccinated:", notVaccinated)

		case 3:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Restarting..")
		}

		fmt.Println("By Karthik Banjan")
	}
}
