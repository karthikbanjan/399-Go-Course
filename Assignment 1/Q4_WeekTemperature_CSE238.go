/*
4. Given a map of day wise temperature in a week.
temp= {"sun":32, "mon":30, "tue":29,"wed":25, "thur":25, "fri":27, "sat":24}
Find which day recorded maximum temperature and find average temperature of the week.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	temperature := map[string]int{"sun": 32, "mon": 30, "tue": 29, "wed": 25, "thur": 25, "fri": 27, "sat": 24}
	max := math.MinInt
	var day string
	sum := 0

	for k, v := range temperature {
		sum += v
		if v > max {
			max = v
			day = k
		}
	}

	fmt.Println("Day with maximum temperature is", day, "with temperature", max, "degree")
	fmt.Println("Average temperature of the week is", sum/len(temperature), "degree")
}
