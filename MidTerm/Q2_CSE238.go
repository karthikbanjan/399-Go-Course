/*

Initialize array of strings which contain phone numbers, names, email ids of employees, like {"joe","98989786",
"abc@yahoo.com",.................}
Extract only email IDs from the array without using in-built functions or strings package.
*/

package main

import "fmt"

func isValidEmail(email string) bool {
	dotPos := -1
	atPos := -1

	for i, v := range email {
		if v == '.' {
			dotPos = i
		} else if v == '@' {
			atPos = i
		}
	}

	if dotPos != -1 && atPos != -1 && dotPos > atPos {
		return true
	}

	return false
}

func main() {
	employeeDetails := [9]string{"Karthik Banjan", "6366805354", "karthikbanjan@protonmail.com",
		"Sam Banjan", "6366805152", "sam@gmail.com",
		"John Banjan", "6366805153", "john@gmail.com"}

	for _, v := range employeeDetails {
		if isValidEmail(v) {
			fmt.Println(v)
		}
	}

}
