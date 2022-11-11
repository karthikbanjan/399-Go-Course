/*
Write go code to print special symbols (@, #, %, &, $ ) from a given string and also
count number of special symbols in the string
*/

package main

import "fmt"

func main() {
	var str string
	println("Enter string: ")
	fmt.Scan(&str)

	var count int
	for _, c := range str {
		if c == '@' || c == '#' || c == '%' || c == '&' || c == '$' {
			fmt.Print(string(c))
			count++
		}
	}
	fmt.Println()
	fmt.Println("Count: ", count)
	fmt.Println("Code by Karthik Banjan")
}
