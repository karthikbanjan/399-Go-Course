/*
5. Initialize a map having key as section and value is another map consists of roll number
and CGPA. Read roll number from the user and find section, and CGPA of the student
map {"cse1”: map[string]float32{"cse0001":8.9, "cse003":7.8}
"cse2": map[string]float32{"cse0010":8.0, "cse0011":7.0}

*/

package main

import "fmt"

func main() {
	db := map[string]map[string]float64{"7CSE5": {"CSE0238": 9.35, "CSE0111": 8.1, "CSE0222": 8.8},
		"7CSE6": {"CSE0239": 8.2, "CSE0112": 8.7, "CSE0224": 9.1}}
	var sec string
	var roll string
	var marks float64 = -1

	fmt.Println("Enter the Roll Number")
	fmt.Scan(&roll)

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
		fmt.Println("Roll number not found!")
	} else {
		fmt.Println("Roll number", roll, "is in section", sec, "and has a CGPA of", marks)
	}
}
