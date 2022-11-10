/*
Write a menu driven program to collect Covid patient details for analysis. Ask these questions to user.
How many times you affected with Covid?  (A for one, B for 2, C for 3 and D for more than 3), store in a list
Ask Covid vaccination status (F for first dose, S for second dose, BS for booster dose, N for no vaccination yet),
store this in another list.
When you affected with Covid last time?
( answers will be before first dose / after first does/after second dose/after booster dose)
Make a  map to store "number of occurrences of Covid" and "when it happened".
Iterate this questions for infinite number of people, make a way to exit program.
*/

package main

import "fmt"

func main() {
	var covidOccur []string
	var vaccStatus []string
	var lastTime []string

	for {
		var co string
		var vs string
		var lt string

		fmt.Print("How many times you affected with Covid?\n(A for one, B for 2, C for 3 and D for more than 3): ")
		fmt.Scan(&co)
		covidOccur = append(covidOccur, co)

		fmt.Print("Covid vaccination status?\n(F for first dose, S for second dose, BS for booster dose, N for no vaccination yet): ")
		fmt.Scan(&vs)
		vaccStatus = append(vaccStatus, vs)

		fmt.Print("When were you affected with Covid last time?\n(beforeFD/afterFD/afterSD/afterBD): ")
		fmt.Scan(&lt)
		lastTime = append(lastTime, lt)

		fmt.Println("Do you want to continue? (Y/N): ")
		var choice string
		fmt.Scan(&choice)
		if choice == "N" {
			break
		}
	}

	occurrenceMap := make(map[string]string)
	for i := 0; i < len(covidOccur); i++ {
		occurrenceMap[covidOccur[i]] = lastTime[i]
	}

	fmt.Println(occurrenceMap)
	fmt.Println("Code by Karthik Banjan")
}
