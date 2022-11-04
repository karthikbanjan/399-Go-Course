// Print words based on character by Karthik Banjan

package main

import "fmt"

func main() {
	var c rune

	fmt.Print("Enter a character: ")
	fmt.Scanf("%c", &c)

	if c == 'A' {
		fmt.Println("Apple")
	} else if c == 'B' {
		fmt.Println("Banana")
	} else if c == 'C' {
		fmt.Println("Coconut")
	} else {
		fmt.Println("Invalid character")
	}
}
