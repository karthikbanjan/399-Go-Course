//Section wise marks with roll number and marks using nested maps

package main

import "fmt"

func main() {
	db := map[string]map[string]int{"7CSE5": {"CSE0238": 100, "CSE0111": 99, "CSE0222": 98},
		"7CSE6": {"CSE0239": 100, "CSE0112": 99, "CSE0224": 98}}
	var sec string
	var roll string
	marks := -1

	fmt.Println("Enter the Roll Number")
	fmt.Scanln(&roll)

	for k, v := range db {
		for k2, v2 := range v {
			if k2 == roll {
				marks = v2
				sec = k
				break
			}
		}
	}

	if marks == -1 {
		fmt.Println("Roll number not found")
	} else {
		fmt.Println("Roll number", roll, "is in section", sec, "and has marks", marks)
	}
}
